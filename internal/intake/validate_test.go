package intake

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestValidate(t *testing.T) {
	cases := []struct {
		name      string
		fixture   string
		wantErr   bool
		wantField string
	}{
		{
			name:    "happy safe intake record",
			fixture: "safe.json",
			wantErr: false,
		},
		{
			name:      "missing license and source rejected",
			fixture:   "malformed_missing_license_source.json",
			wantErr:   true,
			wantField: "source_url",
		},
		{
			name:      "rejected classification without reason rejected",
			fixture:   "rejected_no_reason.json",
			wantErr:   true,
			wantField: "rejection_reason",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := os.ReadFile(filepath.Join("testdata", tc.fixture))
			if err != nil {
				t.Fatalf("read fixture: %v", err)
			}
			rec, err := Validate(data)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil record %+v", rec)
				}
				var verr *Error
				if !errors.As(err, &verr) {
					t.Fatalf("expected *intake.Error, got %T: %v", err, err)
				}
				if tc.wantField != "" && verr.Field != tc.wantField {
					t.Fatalf("expected error field %q, got %q", tc.wantField, verr.Field)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if rec == nil {
				t.Fatal("expected non-nil record")
			}
			if rec.SourceURL == "" || rec.License == "" {
				t.Fatalf("happy record missing required fields: %+v", rec)
			}
			if rec.SafetyClassification == "" {
				t.Fatal("happy record should default to needs-review")
			}
		})
	}
}

func TestValidateRawJSON(t *testing.T) {
	if _, err := Validate([]byte("{not json")); err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}