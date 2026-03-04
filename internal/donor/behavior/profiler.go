// Package behavior derives driver init sequences and access patterns from MMIO traces.
package behavior

import (
	"fmt"
	"strings"
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
)

// Profile is the output of analyzing an MMIO trace.
type Profile struct {
	DeviceBDF    string        `json:"device_bdf"`
	ClassCode    uint32        `json:"class_code"`
	Duration     time.Duration `json:"duration"`
	InitSequence []InitStep    `json:"init_sequence"`
	AccessStats  AccessStats   `json:"access_stats"`
	Timestamp    time.Time     `json:"timestamp"`
}

// InitStep is one reg access during driver init (first unique touch per offset).
type InitStep struct {
	Order     int           `json:"order"`     // step number (1-based)
	Offset    uint32        `json:"offset"`    // BAR register offset
	Type      string        `json:"type"`      // "read" or "write"
	Value     uint32        `json:"value"`     // value read or written
	Timestamp time.Duration `json:"timestamp"` // time since start
	Purpose   string        `json:"purpose"`   // inferred purpose (if known)
}

// AccessStats tracks read/write counts and hot registers.
type AccessStats struct {
	TotalReads    int           `json:"total_reads"`
	TotalWrites   int           `json:"total_writes"`
	UniqueOffsets int           `json:"unique_offsets"`
	HotReads      []OffsetCount `json:"hot_reads"`  // most-read registers
	HotWrites     []OffsetCount `json:"hot_writes"` // most-written registers
}

// OffsetCount is a (register offset, hit count) pair.
type OffsetCount struct {
	Offset uint32 `json:"offset"`
	Count  int    `json:"count"`
}

// --- trace analysis ---

// FromMMIOTrace turns raw MMIO records into an init sequence + stats.
func FromMMIOTrace(trace *mmio.TraceResult, classCode uint32) *Profile {
	if trace == nil {
		return &Profile{Timestamp: time.Now()}
	}

	pattern := mmio.Analyze(trace)
	profile := &Profile{
		DeviceBDF: trace.BDF,
		ClassCode: classCode,
		Duration:  trace.Duration,
		Timestamp: time.Now(),
	}

	// init sequence = first unique access per offset
	seen := make(map[uint32]bool)
	order := 1
	for _, rec := range trace.Records {
		if seen[rec.Offset] {
			continue
		}
		seen[rec.Offset] = true

		step := InitStep{
			Order:     order,
			Offset:    rec.Offset,
			Value:     rec.Value,
			Timestamp: rec.Timestamp,
			Purpose:   inferPurpose(rec.Offset, rec.Type, classCode),
		}
		if rec.Type == mmio.AccessRead {
			step.Type = "read"
		} else {
			step.Type = "write"
		}
		profile.InitSequence = append(profile.InitSequence, step)
		order++
		if order > 64 {
			break
		}
	}

	// aggregate stats
	profile.AccessStats = AccessStats{
		TotalReads:    pattern.TotalReads,
		TotalWrites:   pattern.TotalWrites,
		UniqueOffsets: len(pattern.HotRegisters),
	}
	for _, hr := range pattern.HotRegisters {
		if hr.ReadCount > 0 {
			profile.AccessStats.HotReads = append(profile.AccessStats.HotReads,
				OffsetCount{Offset: hr.Offset, Count: hr.ReadCount})
		}
		if hr.WriteCount > 0 {
			profile.AccessStats.HotWrites = append(profile.AccessStats.HotWrites,
				OffsetCount{Offset: hr.Offset, Count: hr.WriteCount})
		}
	}
	// keep top 10
	if len(profile.AccessStats.HotReads) > 10 {
		profile.AccessStats.HotReads = profile.AccessStats.HotReads[:10]
	}
	if len(profile.AccessStats.HotWrites) > 10 {
		profile.AccessStats.HotWrites = profile.AccessStats.HotWrites[:10]
	}

	return profile
}

// inferPurpose maps well-known BAR offsets to human names, per device class.
func inferPurpose(offset uint32, accessType mmio.AccessType, classCode uint32) string {
	baseClass := (classCode >> 16) & 0xFF
	subClass := (classCode >> 8) & 0xFF

	// NVMe
	if baseClass == 0x01 && subClass == 0x08 {
		switch offset {
		case 0x00:
			return "NVMe CAP (Controller Capabilities)"
		case 0x08:
			return "NVMe VS (Version)"
		case 0x14:
			if accessType == mmio.AccessWrite {
				return "NVMe CC (Controller Configuration — enable)"
			}
			return "NVMe CC (read current config)"
		case 0x1C:
			return "NVMe CSTS (Controller Status — poll RDY)"
		case 0x24:
			return "NVMe AQA (Admin Queue Attributes)"
		}
	}

	// xHCI
	if baseClass == 0x0C && subClass == 0x03 {
		switch offset {
		case 0x00:
			return "xHCI CAPLENGTH/HCIVERSION"
		case 0x04:
			return "xHCI HCSPARAMS1 (structural)"
		case 0x20:
			return "xHCI USBCMD (run/stop)"
		case 0x24:
			return "xHCI USBSTS (status)"
		}
	}

	return ""
}

// --- report ---

// FormatReport prints a short summary of the profile.
func FormatReport(profile *Profile) string {
	if profile == nil {
		return "No behavior profile available.\n"
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("=== Behavior Profile: %s (class 0x%06X) ===\n",
		profile.DeviceBDF, profile.ClassCode))
	sb.WriteString(fmt.Sprintf("Duration: %v | Reads: %d | Writes: %d | Unique regs: %d\n\n",
		profile.Duration, profile.AccessStats.TotalReads,
		profile.AccessStats.TotalWrites, profile.AccessStats.UniqueOffsets))

	if len(profile.InitSequence) > 0 {
		sb.WriteString("--- Initialization Sequence ---\n")
		for _, step := range profile.InitSequence {
			purpose := ""
			if step.Purpose != "" {
				purpose = " ← " + step.Purpose
			}
			sb.WriteString(fmt.Sprintf("  %2d. [%v] %s 0x%03X = 0x%08X%s\n",
				step.Order, step.Timestamp, step.Type, step.Offset, step.Value, purpose))
		}
	}

	return sb.String()
}
