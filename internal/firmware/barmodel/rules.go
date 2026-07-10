package barmodel

import (
	"fmt"
	"sort"

	"github.com/sercanarga/pcileechgen/internal/donor/behavior"
)

func ApplyBehaviorRules(model *BARModel, set *behavior.RuleSet) (*BARModel, error) {
	if set == nil {
		return model, nil
	}
	if err := behavior.Validate(set); err != nil {
		return nil, err
	}
	if model == nil {
		model = &BARModel{Size: set.BARSize}
	}
	if model.Size < set.BARSize {
		model.Size = set.BARSize
	}
	registers := make(map[uint32]int, len(model.Registers))
	for i := range model.Registers {
		registers[model.Registers[i].Offset] = i
	}
	initialValues := make(map[uint32]uint32, len(set.InitialRegisters))
	explicitPolicies := make(map[uint32]bool, len(set.InitialRegisters))
	for _, initial := range set.InitialRegisters {
		if initial.Width != 4 || initial.Offset%4 != 0 {
			return nil, fmt.Errorf("initial register at %#x is not supported by 32-bit BAR RTL", initial.Offset)
		}
		initialValues[initial.Offset] = uint32(initial.Value)
		explicitPolicies[initial.Offset] = initial.WritePolicy != nil
		if index, ok := registers[initial.Offset]; ok {
			if model.Registers[index].IsFSMDriven || model.Registers[index].IsRW1C {
				return nil, fmt.Errorf("initial register at %#x conflicts with existing device behavior", initial.Offset)
			}
			model.Registers[index].Reset = uint32(initial.Value)
			if initial.WritePolicy != nil {
				model.Registers[index].RWMask = uint32(initial.WritePolicy.RWMask)
				model.Registers[index].W1CMask = uint32(initial.WritePolicy.W1CMask)
			}
			continue
		}
		registers[initial.Offset] = len(model.Registers)
		register := BARRegister{
			Offset: initial.Offset, Width: 4, Reset: uint32(initial.Value), Name: fmt.Sprintf("BEHAVIOR_%08X", initial.Offset),
		}
		if initial.WritePolicy != nil {
			register.RWMask = uint32(initial.WritePolicy.RWMask)
			register.W1CMask = uint32(initial.WritePolicy.W1CMask)
		}
		model.Registers = append(model.Registers, register)
	}
	owned := make(map[uint32]struct{})
	for _, rule := range set.Rules {
		if rule.Access != behavior.AccessWrite || rule.Width != 4 || rule.Offset%4 != 0 {
			return nil, fmt.Errorf("rule %q is not supported by 32-bit BAR RTL", rule.ID)
		}
		updates := append([]behavior.RegisterUpdate(nil), rule.Updates...)
		for _, event := range rule.DelayedEvents {
			updates = append(updates, event.Updates...)
		}
		for _, update := range updates {
			if update.Width != 4 || update.Offset%4 != 0 {
				return nil, fmt.Errorf("rule %q update at %#x is not supported by 32-bit BAR RTL", rule.ID, update.Offset)
			}
			if index, ok := registers[update.Offset]; ok {
				if _, alreadyOwned := owned[update.Offset]; alreadyOwned {
					continue
				}
				if model.Registers[index].IsFSMDriven || model.Registers[index].IsRW1C {
					return nil, fmt.Errorf("rule %q update at %#x conflicts with existing device behavior", rule.ID, update.Offset)
				}
				if update.Offset == rule.Offset && !explicitPolicies[update.Offset] {
					model.Registers[index].RWMask = ^uint32(0)
					model.Registers[index].W1CMask = 0
				}
				model.Registers[index].Reset = initialValues[update.Offset]
				model.Registers[index].IsFSMDriven = true
				owned[update.Offset] = struct{}{}
				continue
			}
			registers[update.Offset] = len(model.Registers)
			rwMask := uint32(0)
			if update.Offset == rule.Offset {
				rwMask = ^uint32(0)
			}
			model.Registers = append(model.Registers, BARRegister{
				Offset: update.Offset, Width: 4, Reset: initialValues[update.Offset], RWMask: rwMask,
				Name: fmt.Sprintf("BEHAVIOR_%08X", update.Offset), IsFSMDriven: true,
			})
			owned[update.Offset] = struct{}{}
		}
	}
	sort.Slice(model.Registers, func(i, j int) bool { return model.Registers[i].Offset < model.Registers[j].Offset })
	return model, nil
}
