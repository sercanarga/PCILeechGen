package scrub

import (
	"encoding/binary"
	"fmt"
	"log/slog"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/overlay"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// pcieCapSize is the PCIe v2 capability structure size (60 bytes / 0x3C).
const pcieCapSize = 60

// hasPCIeCap returns true if the capability chain contains a PCIe capability.
func hasPCIeCap(caps []pci.Capability) bool {
	for _, c := range caps {
		if c.ID == pci.CapIDPCIExpress {
			return true
		}
	}
	return false
}

// findFreeCapSpace returns a DWORD-aligned offset in 0x40-0xFF with enough
// contiguous free bytes for a capability of the given size. Returns -1 if
// no suitable gap exists.
func findFreeCapSpace(cs *pci.ConfigSpace, caps []pci.Capability, needed int) int {
	used := make([]bool, pci.ConfigSpaceLegacySize)
	// header region is always occupied
	for i := 0; i < 0x40; i++ {
		used[i] = true
	}
	for _, cap := range caps {
		size := capSizeAt(cs, cap.ID, cap.Offset)
		for i := cap.Offset; i < cap.Offset+size && i < pci.ConfigSpaceLegacySize; i++ {
			used[i] = true
		}
	}

	for start := 0x40; start+needed <= pci.ConfigSpaceLegacySize; start += 4 {
		fits := true
		for i := start; i < start+needed; i++ {
			if used[i] {
				fits = false
				break
			}
		}
		if fits {
			return start
		}
	}
	return -1
}

// buildPCIeCapData builds a 60-byte PCIe v2 endpoint capability structure.
// link speed, width, and speed vector are derived from board config so the
// shadow config space matches the Xilinx IP core parameters.
func buildPCIeCapData(b *board.Board) [pcieCapSize]byte {
	var data [pcieCapSize]byte

	maxSpeed := uint8(firmware.LinkSpeedGen2)
	maxWidth := uint8(1)
	if b != nil {
		maxSpeed = b.MaxLinkSpeedOrDefault()
		if b.PCIeLanes > 0 {
			maxWidth = uint8(b.PCIeLanes)
		}
	}

	// cap ID + next pointer (next set later by caller)
	data[0] = pci.CapIDPCIExpress
	data[1] = 0x00 // end of chain (caller overrides if needed)

	// PCIe Capabilities Register (cap+0x02)
	// version=2, device/port type=0 (endpoint), slot_implemented=0
	binary.LittleEndian.PutUint16(data[0x02:], 0x0002)

	// Device Capabilities (cap+0x04)
	// MPS=0 (128B), phantom=0, ext_tag=0
	// L0s acceptable latency=7 (no limit), L1 acceptable latency=7 (no limit)
	// role-based error reporting=1
	devCap := uint32(7<<6) | uint32(7<<9) | uint32(1<<15)
	binary.LittleEndian.PutUint32(data[0x04:], devCap)

	// Device Control (cap+0x08)
	// relaxed ordering=1, MPS=128B, MRRS=512B (2<<12)
	devCtl := uint16(0x0010) | uint16(2<<12)
	binary.LittleEndian.PutUint16(data[0x08:], devCtl)

	// Device Status (cap+0x0A) = 0x0000

	// Link Capabilities (cap+0x0C)
	// max speed, max width, L0s exit latency=6 (1-2us), L1 exit latency=6 (32-64us)
	linkCap := uint32(maxSpeed) |
		(uint32(maxWidth) << 4) |
		(6 << 12) | // L0s exit latency
		(6 << 15)   // L1 exit latency
	binary.LittleEndian.PutUint32(data[0x0C:], linkCap)

	// Link Control (cap+0x10) = 0x0000 (no ASPM)

	// Link Status (cap+0x12)
	// current speed, negotiated width, slot clock config=1
	linkStatus := uint16(maxSpeed) | (uint16(maxWidth) << 4) | (1 << 12)
	binary.LittleEndian.PutUint16(data[0x12:], linkStatus)

	// Slot Capabilities (cap+0x14) = 0 (endpoint, no slot)
	// Slot Control (cap+0x18) = 0
	// Slot Status (cap+0x1A) = 0
	// Root Control (cap+0x1C) = 0
	// Root Capabilities (cap+0x1E) = 0
	// Root Status (cap+0x20) = 0

	// Device Capabilities 2 (cap+0x24) = 0 (clampDeviceCapPass handles it)

	// Device Control 2 (cap+0x28) = 0
	// Device Status 2 (cap+0x2A) = 0

	// Link Capabilities 2 (cap+0x2C)
	// supported link speeds vector (bits 7:1)
	var speedVec uint32
	for s := uint8(1); s <= maxSpeed; s++ {
		speedVec |= 1 << s
	}
	binary.LittleEndian.PutUint32(data[0x2C:], speedVec)

	// Link Control 2 (cap+0x30)
	// target link speed = maxSpeed
	binary.LittleEndian.PutUint16(data[0x30:], uint16(maxSpeed))

	// Link Status 2 (cap+0x32) = 0
	// Slot Capabilities 2 (cap+0x34) = 0
	// Slot Control 2 (cap+0x38) = 0
	// Slot Status 2 (cap+0x3A) = 0

	return data
}

// injectPCIeCapIfMissing adds a synthesized PCIe capability when the donor
// config space lacks one. This is needed for conventional PCI donors
// (e.g. Creative Labs VEN_1102 behind a PCIe-to-PCI bridge) where the
// original device has no PCIe capability structure.
//
// Without this injection, Windows classifies the FPGA as PciConventional
// (DeviceType=0) and fails resource allocation with
// STATUS_DEVICE_ENUMERATION_ERROR (0xC0000366).
func injectPCIeCapIfMissing(cs *pci.ConfigSpace, b *board.Board, om *overlay.Map, ctx *ScrubContext) {
	if hasPCIeCap(ctx.Caps) {
		return
	}

	slog.Info("donor device lacks PCIe capability, injecting synthesized PCIe cap")

	// ensure capabilities list bit is set in Status register
	status := cs.Status()
	if status&0x0010 == 0 {
		om.WriteU16(0x06, status|0x0010, "set Capabilities List bit in Status")
	}

	// ensure CapPtr is valid when there are no existing caps
	if len(ctx.Caps) == 0 && cs.CapabilityPointer() == 0 {
		slog.Info("donor has no capability chain, creating minimal PM + MSI + PCIe chain")
		injectFullCapChain(cs, b, om, ctx)
		return
	}

	capOffset := findFreeCapSpace(cs, ctx.Caps, pcieCapSize)
	if capOffset < 0 {
		slog.Warn("no free space in 0x40-0xFF for PCIe capability injection")
		return
	}

	// build PCIe cap data
	capData := buildPCIeCapData(b)

	// write cap data to config space
	for i := 0; i < pcieCapSize; i++ {
		om.WriteU8(capOffset+i, capData[i],
			fmt.Sprintf("inject PCIe cap byte at 0x%02X", capOffset+i))
	}

	// link into the end of existing capability chain
	lastCap := ctx.Caps[len(ctx.Caps)-1]
	om.WriteU8(lastCap.Offset+1, uint8(capOffset),
		fmt.Sprintf("link cap 0x%02X at 0x%02X -> injected PCIe cap at 0x%02X",
			lastCap.ID, lastCap.Offset, capOffset))

	// re-parse caps so downstream passes see the newly injected cap
	ctx.Caps = pci.ParseCapabilities(cs)

	speedName := firmware.LinkSpeedName(boardMaxSpeed(b))
	lanes := boardLanes(b)
	slog.Info("injected PCIe capability",
		"offset", fmt.Sprintf("0x%02X", capOffset),
		"link_speed", speedName,
		"link_width", fmt.Sprintf("x%d", lanes))
}

// injectFullCapChain creates a minimal PM -> MSI -> PCIe capability chain
// when the donor has no capabilities at all. Starts at offset 0x40.
func injectFullCapChain(cs *pci.ConfigSpace, b *board.Board, om *overlay.Map, ctx *ScrubContext) {
	const (
		pmOffset   = 0x40 // PM cap: 8 bytes (0x40-0x47)
		msiOffset  = 0x48 // MSI cap: 12 bytes 32-bit (0x48-0x53), padded to 0x54
		pcieOffset = 0x54 // PCIe cap: 60 bytes (0x54-0x8F)
	)

	// set CapPtr to start of chain
	om.WriteU8(0x34, pmOffset, "set CapPtr to injected PM cap")

	// PM capability (8 bytes)
	om.WriteU8(pmOffset, pci.CapIDPowerManagement, "inject PM cap ID")
	om.WriteU8(pmOffset+1, msiOffset, "PM next -> MSI")
	// PMC: version 3, no PME support from any D-state.
	// advertising PME from D3hot/D3cold tells Windows the device can wake
	// from D3, which causes the PM framework to aggressively transition to
	// D3hot after ~5 minutes of idle. the FPGA IP core can't properly
	// handle D3 and stops processing TLPs.
	om.WriteU16(pmOffset+2, 0x0003, "PM capabilities (v3, no PME support)")
	// PMCSR: D0 state, NoSoftReset, PME_Enable=0
	om.WriteU16(pmOffset+4, 0x0008, "PMCSR: D0, NoSoftReset, PME disabled")
	om.WriteU16(pmOffset+6, 0x0000, "PM bridge/data")

	// MSI capability (10 bytes, 32-bit address, 1 vector)
	om.WriteU8(msiOffset, pci.CapIDMSI, "inject MSI cap ID")
	om.WriteU8(msiOffset+1, pcieOffset, "MSI next -> PCIe")
	// Message Control: 32-bit, 1 vector requested, disabled
	om.WriteU16(msiOffset+2, 0x0000, "MSI message control (32-bit, 1 vec)")
	om.WriteU32(msiOffset+4, 0x00000000, "MSI message address")
	om.WriteU16(msiOffset+8, 0x0000, "MSI message data")
	// pad to DWORD boundary
	om.WriteU16(msiOffset+10, 0x0000, "MSI padding")

	// PCIe capability (60 bytes)
	capData := buildPCIeCapData(b)
	for i := 0; i < pcieCapSize; i++ {
		om.WriteU8(pcieOffset+i, capData[i],
			fmt.Sprintf("inject PCIe cap byte at 0x%02X", pcieOffset+i))
	}

	// re-parse caps so downstream passes see the full chain
	ctx.Caps = pci.ParseCapabilities(cs)

	slog.Info("injected full capability chain (PM + MSI + PCIe)",
		"pm_offset", fmt.Sprintf("0x%02X", pmOffset),
		"msi_offset", fmt.Sprintf("0x%02X", msiOffset),
		"pcie_offset", fmt.Sprintf("0x%02X", pcieOffset),
		"link_speed", firmware.LinkSpeedName(boardMaxSpeed(b)),
		"link_width", fmt.Sprintf("x%d", boardLanes(b)))
}

// boardMaxSpeed returns the board's max link speed, defaulting to Gen2 if nil.
func boardMaxSpeed(b *board.Board) uint8 {
	if b != nil {
		return b.MaxLinkSpeedOrDefault()
	}
	return firmware.LinkSpeedGen2
}

// boardLanes returns the board's PCIe lane count, defaulting to 1 if nil.
func boardLanes(b *board.Board) int {
	if b != nil && b.PCIeLanes > 0 {
		return b.PCIeLanes
	}
	return 1
}
