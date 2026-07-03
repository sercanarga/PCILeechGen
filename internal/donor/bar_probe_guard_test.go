package donor

import "testing"

func TestShouldProbeBAR(t *testing.T) {
	good := []byte{0x00, 0x01, 0x02, 0x03}
	ff := []byte{0xFF, 0xFF, 0xFF, 0xFF}

	cases := []struct {
		name      string
		enabled   bool
		classCode uint32
		content   []byte
		want      bool
	}{
		{"disabled", false, 0x010802, good, false},
		{"network ethernet skipped", true, 0x020000, good, false},
		{"network other-subclass skipped (aquantia 10GbE)", true, 0x028000, good, false},
		{"all-FF skipped", true, 0x010802, ff, false},
		{"nvme probed", true, 0x010802, good, true},
		{"audio probed", true, 0x040300, good, true},
		{"no content still probed (non-network)", true, 0x010802, nil, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := shouldProbeBAR(c.enabled, c.classCode, c.content); got != c.want {
				t.Errorf("shouldProbeBAR(%v, %#x, len=%d) = %v, want %v", c.enabled, c.classCode, len(c.content), got, c.want)
			}
		})
	}
}
