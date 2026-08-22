package graph

import (
	"fmt"
)

// CheckAcyclic checks if the knowledge graph is acyclic using DFS
func CheckAcyclic(guides map[string]Guide) error {
	const (
		unvisited = 0
		visiting  = 1
		visited   = 2
	)

	state := make(map[string]int, len(guides))

	var dfs func(id string) error
	dfs = func(id string) error {
		state[id] = visiting

		if guide, exists := guides[id]; exists {
			neighbors := append([]string(nil), guide.Metadata.Prerequisites...)
			neighbors = append(neighbors, guide.Metadata.SubGuides...)

			for _, neighborID := range neighbors {
				if _, exists := guides[neighborID]; !exists {
					continue
				}
				if state[neighborID] == visiting {
					return fmt.Errorf("cycle detected in graph involving guide '%s'", neighborID)
				}
				if state[neighborID] == unvisited {
					if err := dfs(neighborID); err != nil {
						return err
					}
				}
			}
		}

		state[id] = visited
		return nil
	}

	for id := range guides {
		if state[id] == unvisited {
			if err := dfs(id); err != nil {
				return err
			}
		}
	}

	return nil
}

// getScopeValue helper returns the hierarchical index of the given scope (1-indexed)
func getScopeValue(scope string, config QuestConfig) int {
	for i, s := range config.Scopes {
		if s == scope {
			return i + 1
		}
	}
	return 0
}

func isValidClarity(clarity string, config QuestConfig) bool {
	for _, c := range config.Clarities {
		if c == clarity {
			return true
		}
	}
	return false
}

// ValidateGraph checks structural constraints of the knowledge graph
func ValidateGraph(guides map[string]Guide, config QuestConfig) error {
	for _, guide := range guides {
		// Validate Scope
		if getScopeValue(guide.Metadata.Scope, config) == 0 {
			return fmt.Errorf("guide '%s' has invalid scope: '%s'", guide.ID, guide.Metadata.Scope)
		}

		// Validate Clarity
		if !isValidClarity(guide.Metadata.Clarity, config) {
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
			if getScopeValue(sub.Metadata.Scope, config) >= getScopeValue(guide.Metadata.Scope, config) {
				return fmt.Errorf("guide '%s' (scope: '%s') has sub_guide '%s' with invalid scope: '%s'. Sub-guides must have a strictly smaller scope",
					guide.ID, guide.Metadata.Scope, subID, sub.Metadata.Scope)
			}
		}
	}
	return CheckAcyclic(guides)
}
