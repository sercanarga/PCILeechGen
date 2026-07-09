package devicemodel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type modelWire Model

func (m Model) MarshalJSON() ([]byte, error) {
	if err := (&m).Validate(); err != nil {
		return nil, fmt.Errorf("marshal device model: %w", err)
	}
	return json.Marshal(modelWire(m))
}

func (m *Model) UnmarshalJSON(data []byte) error {
	if m == nil {
		return fmt.Errorf("unmarshal device model: nil destination")
	}
	var decoded modelWire
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&decoded); err != nil {
		return fmt.Errorf("unmarshal device model: %w", err)
	}
	if err := requireJSONEOF(decoder); err != nil {
		return fmt.Errorf("unmarshal device model: %w", err)
	}
	candidate := Model(decoded)
	if err := candidate.Validate(); err != nil {
		return fmt.Errorf("unmarshal device model: %w", err)
	}
	*m = candidate
	return nil
}

func (m *Model) ToJSON() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("marshal device model: nil model")
	}
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(data, '\n'), nil
}

func ParseJSON(data []byte) (*Model, error) {
	var model Model
	if err := json.Unmarshal(data, &model); err != nil {
		return nil, err
	}
	return &model, nil
}

func FromJSON(data []byte) (*Model, error) { return ParseJSON(data) }

func requireJSONEOF(decoder *json.Decoder) error {
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		if err == nil {
			return fmt.Errorf("unexpected trailing JSON value")
		}
		return err
	}
	return nil
}
