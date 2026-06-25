package output

import (
	"encoding/json"
	"fmt"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware/barprofile"
)

func (ow *OutputWriter) writeBARBehaviorProfile(ctx *donor.DeviceContext) error {
	profile := barprofile.Build(ctx)
	data, err := json.MarshalIndent(profile, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal BAR behavior profile: %w", err)
	}
	if err := ow.writeFile("bar_behavior_profile.json", string(data)+"\n"); err != nil {
		return fmt.Errorf("failed to write BAR behavior profile: %w", err)
	}
	return nil
}
