package tclgen

import "testing"

func TestLinkSpeedToTrgt_AllSpeeds(t *testing.T) {
	tests := []struct {
		speed uint8
		want  string
	}{
		{1, "4'h1"}, // Gen1
		{2, "4'h2"}, // Gen2
		{3, "4'h3"}, // Gen3
		{0, "4'h2"}, // default
	}
	for _, tt := range tests {
		got := linkSpeedToTrgt(tt.speed)
		if got != tt.want {
			t.Errorf("linkSpeedToTrgt(%d) = %q, want %q", tt.speed, got, tt.want)
		}
	}
}

func TestLinkWidthToTCL_AllWidths(t *testing.T) {
	tests := []struct {
		width uint8
		want  string
	}{
		{1, "X1"},
		{2, "X2"},
		{4, "X4"},
		{8, "X8"},
		{0, "X1"},  // default
		{16, "X1"}, // unknown → default
	}
	for _, tt := range tests {
		got := linkWidthToTCL(tt.width)
		if got != tt.want {
			t.Errorf("linkWidthToTCL(%d) = %q, want %q", tt.width, got, tt.want)
		}
	}
}

func TestBarSizeToTCL_SmallSizes(t *testing.T) {
	// 1KB should clamp to 4KB minimum
	scale, size := barSizeToTCL(1024)
	if scale != "Kilobytes" || size != "4" {
		t.Errorf("1KB: got %s/%s, want Kilobytes/4", scale, size)
	}

	// 64KB
	scale, size = barSizeToTCL(64 * 1024)
	if scale != "Kilobytes" || size != "64" {
		t.Errorf("64KB: got %s/%s, want Kilobytes/64", scale, size)
	}

	// 2MB
	scale, size = barSizeToTCL(2 * 1024 * 1024)
	if scale != "Megabytes" || size != "2" {
		t.Errorf("2MB: got %s/%s, want Megabytes/2", scale, size)
	}
}
