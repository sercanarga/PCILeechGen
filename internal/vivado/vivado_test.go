package vivado

import (
	"testing"
)

func TestFindValidation(t *testing.T) {
	// Non-existent path should fail
	_, err := Find("/nonexistent/path/to/vivado")
	if err == nil {
		t.Error("Find should fail for non-existent custom path")
	}
}

func TestFindNoArgs(t *testing.T) {
	// Without vivado installed, Find("") should fail gracefully
	_, err := Find("")
	if err == nil {
		// Vivado is actually installed - still valid
		return
	}
	if err.Error() == "" {
		t.Error("Error message should not be empty")
	}
}

func TestVivadoBinaryPath(t *testing.T) {
	v := &Vivado{
		Path:    "/tools/Xilinx/Vivado/2022.2",
		Version: "2022.2",
	}

	path := v.BinaryPath()
	expected := "/tools/Xilinx/Vivado/2022.2/bin/vivado"
	if path != expected {
		t.Errorf("BinaryPath() = %q, want %q", path, expected)
	}
}

func TestBuilderDefaults(t *testing.T) {
	b, _ := Find("") // will fail but ok for testing builder creation
	_ = b

	opts := BuildOptions{}
	builder := NewBuilder(nil, opts)

	if builder.opts.Jobs != 4 {
		t.Errorf("Default jobs = %d, want 4", builder.opts.Jobs)
	}
	if builder.opts.Timeout != 3600 {
		t.Errorf("Default timeout = %d, want 3600", builder.opts.Timeout)
	}
	if builder.opts.OutputDir != "pcileech_datastore" {
		t.Errorf("Default output = %q, want 'pcileech_datastore'", builder.opts.OutputDir)
	}
}
