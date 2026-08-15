package graph

import (
	"fmt"
)

// ValidateGraph checks structural constraints of the knowledge graph
func ValidateGraph(guides map[string]Guide) error {
	for _, guide := range guides {
		// Validate Scope
		if ScopeValue(guide.Metadata.Scope) == 0 {
			return fmt.Errorf("guide '%s' has invalid scope: '%s'", guide.ID, guide.Metadata.Scope)
		}

		// Validate Clarity
		switch guide.Metadata.Clarity {
		case ClarityStrict, ClarityDetailed, ClarityIntroductory, ClarityVague:
			// valid
		default:
			return fmt.Errorf("guide '%s' has invalid clarity: '%s'", guide.ID, guide.Metadata.Clarity)
		}

		// Validate Prerequisites
		for _, prereqID := range guide.Metadata.Prerequisites {
			prereq, exists := guides[prereqID]
			if !exists {
				return fmt.Errorf("guide '%s' references unknown prerequisite: '%s'", guide.ID, prereqID)
			}
			if prereq.Metadata.Scope != guide.Metadata.Scope {
				return fmt.Errorf("guide '%s' (scope: '%s') has prerequisite '%s' with mismatched scope: '%s'. Horizontal edges must have exactly identical scope",
					guide.ID, guide.Metadata.Scope, prereqID, prereq.Metadata.Scope)
			}
		}

		// Validate SubGuides
		for _, subID := range guide.Metadata.SubGuides {
			sub, exists := guides[subID]
			if !exists {
				return fmt.Errorf("guide '%s' references unknown sub_guide: '%s'", guide.ID, subID)
			}
			if ScopeValue(sub.Metadata.Scope) >= ScopeValue(guide.Metadata.Scope) {
				return fmt.Errorf("guide '%s' (scope: '%s') has sub_guide '%s' with invalid scope: '%s'. Sub-guides must have a strictly smaller scope",
					guide.ID, guide.Metadata.Scope, subID, sub.Metadata.Scope)
			}
		}
	}
	return nil
}
