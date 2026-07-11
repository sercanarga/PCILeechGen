package main

import (
	"fmt"
	"io"
	"math/bits"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware/codegen"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

type donorReport struct {
	Identity     donorReportIdentity     `json:"identity"`
	ConfigSpace  donorReportConfigSpace  `json:"config_space"`
	Capabilities []donorReportCapability `json:"capabilities"`
	BARs         []donorReportBAR        `json:"bars"`
	Issues       []string                `json:"issues,omitempty"`
}

type donorReportIdentity struct {
	BDF       string `json:"bdf"`
	VendorID  string `json:"vendor_id"`
	DeviceID  string `json:"device_id"`
	Revision  string `json:"revision"`
	ClassCode string `json:"class_code"`
	Class     string `json:"class"`
}

type donorReportConfigSpace struct {
	Bytes        int `json:"bytes"`
	WritableBits int `json:"writable_bits"`
	ReadOnlyBits int `json:"read_only_bits"`
}

type donorReportCapability struct {
	Kind    string `json:"kind"`
	ID      string `json:"id"`
	Name    string `json:"name"`
	Offset  int    `json:"offset"`
	Version uint8  `json:"version,omitempty"`
}

type donorReportBAR struct {
	Index        int    `json:"index"`
	Type         string `json:"type"`
	Size         uint64 `json:"size"`
	Prefetchable bool   `json:"prefetchable"`
	Is64Bit      bool   `json:"is_64bit"`
}

func buildDonorReport(ctx *donor.DeviceContext) donorReport {
	r := donorReport{
		Identity: donorReportIdentity{
			BDF: ctx.Device.BDF.String(), VendorID: fmt.Sprintf("%04x", ctx.Device.VendorID),
			DeviceID: fmt.Sprintf("%04x", ctx.Device.DeviceID), Revision: fmt.Sprintf("%02x", ctx.Device.RevisionID),
			ClassCode: fmt.Sprintf("%06x", ctx.Device.ClassCode), Class: ctx.Device.ClassDescription(),
		},
	}
	if ctx.ConfigSpace == nil {
		r.Issues = []string{"configuration space is nil"}
		return r
	}

	masks := codegen.GenerateWritemask(ctx.ConfigSpace)
	for _, mask := range masks {
		r.ConfigSpace.WritableBits += bits.OnesCount32(mask)
	}
	r.ConfigSpace.Bytes = ctx.ConfigSpace.Size
	r.ConfigSpace.ReadOnlyBits = ctx.ConfigSpace.Size*8 - r.ConfigSpace.WritableBits
	r.Issues = append(pci.ValidateCapabilityChains(ctx.ConfigSpace), donor.ValidateDeviceLayout(ctx)...)

	for _, cap := range ctx.Capabilities {
		r.Capabilities = append(r.Capabilities, donorReportCapability{Kind: "standard", ID: fmt.Sprintf("%02x", cap.ID), Name: pci.CapabilityName(cap.ID), Offset: cap.Offset})
	}
	for _, cap := range ctx.ExtCapabilities {
		r.Capabilities = append(r.Capabilities, donorReportCapability{Kind: "extended", ID: fmt.Sprintf("%04x", cap.ID), Name: pci.ExtCapabilityName(cap.ID), Offset: cap.Offset, Version: cap.Version})
	}
	for _, bar := range ctx.BARs {
		r.BARs = append(r.BARs, donorReportBAR{Index: bar.Index, Type: bar.Type, Size: bar.Size, Prefetchable: bar.Prefetchable, Is64Bit: bar.Is64Bit})
	}
	return r
}

func printDonorReport(w io.Writer, r donorReport) {
	fmt.Fprintln(w, "Donor analysis")
	fmt.Fprintf(w, "  Identity:       %s:%s rev %s, class %s (%s)\n", r.Identity.VendorID, r.Identity.DeviceID, r.Identity.Revision, r.Identity.ClassCode, r.Identity.Class)
	fmt.Fprintf(w, "  Config space:   %d bytes\n", r.ConfigSpace.Bytes)
	fmt.Fprintf(w, "  Writable bits:  %d\n", r.ConfigSpace.WritableBits)
	fmt.Fprintf(w, "  Read-only bits: %d\n", r.ConfigSpace.ReadOnlyBits)
	for _, cap := range r.Capabilities {
		fmt.Fprintf(w, "  Capability:     0x%03x %s (%s)\n", cap.Offset, cap.Name, cap.Kind)
	}
	for _, bar := range r.BARs {
		fmt.Fprintf(w, "  BAR%d:           %s, %d bytes, 64-bit=%t, prefetchable=%t\n", bar.Index, bar.Type, bar.Size, bar.Is64Bit, bar.Prefetchable)
	}
	for _, issue := range r.Issues {
		fmt.Fprintf(w, "  Issue:          %s\n", issue)
	}
	fmt.Fprintln(w)
}
