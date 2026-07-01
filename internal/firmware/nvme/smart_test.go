package nvme

import "testing"

func TestBuildSMART_NonZeroPlausible(t *testing.T) {
	s := BuildSMART()
	if s == nil {
		t.Fatal("BuildSMART returned nil")
	}
	if s.PowerOnHours < 100 {
		t.Errorf("PowerOnHours too low for a used drive: %d", s.PowerOnHours)
	}
	if s.PowerCycles < 10 {
		t.Errorf("PowerCycles too low: %d", s.PowerCycles)
	}
	if s.DataUnitsWritten == 0 {
		t.Error("DataUnitsWritten should be non-zero for a used drive")
	}
	if s.DataUnitsRead == 0 {
		t.Error("DataUnitsRead should be non-zero for a used drive")
	}
	if s.HostReadCommands == 0 || s.HostWriteCommands == 0 {
		t.Error("host read/write commands should be non-zero")
	}
}

// TestBuildSMART_ScalesWithPowerOnHours verifies data-unit and command counts
// correlate with power-on hours so the clone looks consistently worn.
func TestBuildSMART_ScalesWithPowerOnHours(t *testing.T) {
	for i := 0; i < 50; i++ {
		s := BuildSMART()
		if s.DataUnitsWritten < uint64(s.PowerOnHours) {
			t.Fatalf("DataUnitsWritten (%d) should scale with PowerOnHours (%d)", s.DataUnitsWritten, s.PowerOnHours)
		}
	}
}

// TestBuildSMART_BoundedPlausible checks ranges hold over many draws.
func TestBuildSMART_BoundedPlausible(t *testing.T) {
	for i := 0; i < 100; i++ {
		s := BuildSMART()
		if s.PowerOnHours > 20000 {
			t.Errorf("PowerOnHours implausibly high: %d", s.PowerOnHours)
		}
		if s.PowerCycles > 2000 {
			t.Errorf("PowerCycles implausibly high: %d", s.PowerCycles)
		}
		if s.UnsafeShutdowns > 50 {
			t.Errorf("UnsafeShutdowns implausibly high: %d", s.UnsafeShutdowns)
		}
		if s.MediaErrors > 10 {
			t.Errorf("MediaErrors implausibly high: %d", s.MediaErrors)
		}
		if s.ErrorLogEntries > 50 {
			t.Errorf("ErrorLogEntries implausibly high: %d", s.ErrorLogEntries)
		}
		// 48-bit NVMe log fields must fit.
		if s.DataUnitsRead>>48 != 0 || s.DataUnitsWritten>>48 != 0 {
			t.Errorf("data-unit counters exceed 48 bits")
		}
	}
}
