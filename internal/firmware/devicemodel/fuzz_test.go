package devicemodel

import "testing"

// FuzzFromJSONModel exercises schema, BAR topology, reset-image, and MSI-X
// validation with mutations of a complete canonical model.
func FuzzFromJSONModel(f *testing.F) {
	seed, err := validTestModel().ToJSON()
	if err != nil {
		f.Fatalf("marshal seed model: %v", err)
	}
	f.Add(seed)
	f.Add([]byte(`{"schema_version":1,"name":"broken"}`))
	f.Add([]byte{0xff, 0x00, 0x7f})
	f.Fuzz(func(t *testing.T, data []byte) {
		model, err := FromJSON(data)
		if err == nil && model == nil {
			t.Fatal("FromJSON returned nil model without an error")
		}
	})
}
