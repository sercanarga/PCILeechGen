// Package intake validates external adaptation intake records that describe
// feature or profile ideas proposed for adoption into PCILeechGen.
//
// The validator enforces only structural completeness: required provenance
// fields are present and rejection records carry a reason. It does NOT bless
// any external code or profile as safe. Promoting a record to
// safety_classification "safe" is a human act recorded outside this schema.
package intake

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Safety classifications supported by the intake schema.
const (
	Safe        = "safe"
	NeedsReview = "needs-review"
	Rejected    = "rejected"
)

// Record is the structured intake metadata for an external feature/profile idea.
// See docs/external-adaptation-policy.md for the full schema and guardrails.
type Record struct {
	SourceURL            string   `json:"source_url"`
	License              string   `json:"license"`
	CommitRef            string   `json:"commit_ref,omitempty"`
	Scope                string   `json:"scope,omitempty"`
	SafetyClassification string   `json:"safety_classification,omitempty"`
	ImportedArtifacts    []string `json:"imported_artifacts,omitempty"`
	Attribution          string   `json:"attribution,omitempty"`
	RejectionReason      string   `json:"rejection_reason,omitempty"`
}

// Error is a typed validation error returned by Validate.
type Error struct {
	Field   string
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("intake: %s: %s", e.Field, e.Message)
}

// Validate parses and structurally validates an intake record JSON blob.
// It returns nil for a structurally complete record. A nil result does NOT
// mean the record is safe — only that its provenance metadata is present.
func Validate(data []byte) (*Record, error) {
	var r Record
	if err := json.Unmarshal(data, &r); err != nil {
		return nil, &Error{Field: "json", Message: fmt.Sprintf("invalid JSON: %v", err)}
	}

	var errs []*Error
	if strings.TrimSpace(r.SourceURL) == "" {
		errs = append(errs, &Error{Field: "source_url", Message: "required field is missing"})
	}
	if strings.TrimSpace(r.License) == "" {
		errs = append(errs, &Error{Field: "license", Message: "required field is missing"})
	}
	if r.SafetyClassification == "" {
		r.SafetyClassification = NeedsReview
	}
	if r.SafetyClassification == Rejected && strings.TrimSpace(r.RejectionReason) == "" {
		errs = append(errs, &Error{Field: "rejection_reason", Message: "required when safety_classification is rejected"})
	}

	// ponytail: ceiling — only structural checks. No banned-keyword scanning,
	// no source-domain registry, no signatures. Upgrade path: add a controlled
	// vocabulary for `scope` and a banned-source registry once intake volume
	// justifies it.
	if len(errs) > 0 {
		// Return the first error as the typed error; the rest are reachable
		// via the validation transcript in tests. Keeping one error keeps the
		// API tiny.
		return nil, errs[0]
	}
	return &r, nil
}