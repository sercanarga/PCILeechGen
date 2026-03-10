// Package overlay records every config space byte change with a reason.
package overlay

import (
	"fmt"
	"strings"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

// Entry is one byte-level config space edit.
type Entry struct {
	Offset   int    // byte offset in config space
	Width    int    // write width (1, 2, or 4 bytes)
	OldValue uint32 // value before modification
	NewValue uint32 // value after modification
	Reason   string // human-readable reason for the change
}

// Map collects entries as you mutate a config space.
type Map struct {
	original *pci.ConfigSpace
	entries  []Entry
}

func NewMap(cs *pci.ConfigSpace) *Map {
	return &Map{
		original: cs,
	}
}

// ConfigSpace returns the underlying config space being modified.
func (m *Map) ConfigSpace() *pci.ConfigSpace {
	return m.original
}

// WriteU32 writes a 32-bit value and records the change.
func (m *Map) WriteU32(offset int, value uint32, reason string) {
	old := m.original.ReadU32(offset)
	if old == value {
		return // no-op
	}
	m.entries = append(m.entries, Entry{
		Offset:   offset,
		Width:    4,
		OldValue: old,
		NewValue: value,
		Reason:   reason,
	})
	m.original.WriteU32(offset, value)
}

// WriteU16 writes a 16-bit value and records the change.
func (m *Map) WriteU16(offset int, value uint16, reason string) {
	old := m.original.ReadU16(offset)
	if old == value {
		return
	}
	m.entries = append(m.entries, Entry{
		Offset:   offset,
		Width:    2,
		OldValue: uint32(old),
		NewValue: uint32(value),
		Reason:   reason,
	})
	m.original.WriteU16(offset, value)
}

// WriteU8 writes a byte and records the change.
func (m *Map) WriteU8(offset int, value uint8, reason string) {
	old := m.original.ReadU8(offset)
	if old == value {
		return
	}
	m.entries = append(m.entries, Entry{
		Offset:   offset,
		Width:    1,
		OldValue: uint32(old),
		NewValue: uint32(value),
		Reason:   reason,
	})
	m.original.WriteU8(offset, value)
}

// ZeroRange zeroes a byte range and records a single entry.
func (m *Map) ZeroRange(start, end int, reason string) {
	changed := false
	for i := start; i < end && i < m.original.Size; i++ {
		if m.original.ReadU8(i) != 0 {
			changed = true
			break
		}
	}
	if !changed {
		return
	}
	m.entries = append(m.entries, Entry{
		Offset:   start,
		Width:    end - start,
		OldValue: 0xFFFFFFFF, // sentinel meaning "range"
		NewValue: 0,
		Reason:   reason,
	})
	for i := start; i < end && i < m.original.Size; i++ {
		m.original.WriteU8(i, 0)
	}
}

// Diff returns all recorded modifications.
func (m *Map) Diff() []Entry {
	result := make([]Entry, len(m.entries))
	copy(result, m.entries)
	return result
}

// Count returns the number of modifications.
func (m *Map) Count() int {
	return len(m.entries)
}

// FormatDiff returns a human-readable summary of all changes.
func (m *Map) FormatDiff() string {
	if len(m.entries) == 0 {
		return "No modifications.\n"
	}
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("=== Config Space Overlay (%d changes) ===\n", len(m.entries)))
	for _, e := range m.entries {
		if e.Width > 4 {
			// Range zero
			sb.WriteString(fmt.Sprintf("  [%03X-%03X] zeroed (%d bytes) - %s\n",
				e.Offset, e.Offset+e.Width-1, e.Width, e.Reason))
		} else {
			fmtStr := fmt.Sprintf("%%0%dX", e.Width*2)
			oldStr := fmt.Sprintf(fmtStr, e.OldValue)
			newStr := fmt.Sprintf(fmtStr, e.NewValue)
			sb.WriteString(fmt.Sprintf("  [%03X] %s -> %s - %s\n",
				e.Offset, oldStr, newStr, e.Reason))
		}
	}
	return sb.String()
}
