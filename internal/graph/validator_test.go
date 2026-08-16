package graph

import (
	"testing"
)

func TestValidateGraph(t *testing.T) {
	tests := []struct {
		name    string
		guides  map[string]Guide
		wantErr bool
	}{
		{
			name: "valid single guide",
			guides: map[string]Guide{
				"guide1": {
					ID: "guide1",
					Metadata: GuideMetadata{
						Scope:   ScopeDefinition,
						Clarity: ClarityStrict,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid scope",
			guides: map[string]Guide{
				"guide1": {
					ID: "guide1",
					Metadata: GuideMetadata{
						Scope:   "invalid_scope",
						Clarity: ClarityStrict,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "invalid clarity",
			guides: map[string]Guide{
				"guide1": {
					ID: "guide1",
					Metadata: GuideMetadata{
						Scope:   ScopeDefinition,
						Clarity: "invalid_clarity",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "valid prerequisite",
			guides: map[string]Guide{
				"guide1": {
					ID: "guide1",
					Metadata: GuideMetadata{
						Scope:         ScopeLesson,
						Clarity:       ClarityDetailed,
						Prerequisites: []string{"guide2"},
					},
				},
				"guide2": {
					ID: "guide2",
					Metadata: GuideMetadata{
						Scope:   ScopeLesson,
						Clarity: ClarityDetailed,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "unknown prerequisite",
			guides: map[string]Guide{
				"guide1": {
					ID: "guide1",
					Metadata: GuideMetadata{
						Scope:         ScopeLesson,
						Clarity:       ClarityDetailed,
						Prerequisites: []string{"unknown_guide"},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "mismatched prerequisite scope",
			guides: map[string]Guide{
				"guide1": {
					ID: "guide1",
					Metadata: GuideMetadata{
						Scope:         ScopeLesson,
						Clarity:       ClarityDetailed,
						Prerequisites: []string{"guide2"},
					},
				},
				"guide2": {
					ID: "guide2",
					Metadata: GuideMetadata{
						Scope:   ScopeDefinition,
						Clarity: ClarityDetailed,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "valid sub-guide",
			guides: map[string]Guide{
				"guide1": {
					ID: "guide1",
					Metadata: GuideMetadata{
						Scope:     ScopeLesson,
						Clarity:   ClarityDetailed,
						SubGuides: []string{"guide2"},
					},
				},
				"guide2": {
					ID: "guide2",
					Metadata: GuideMetadata{
						Scope:   ScopeExplanation,
						Clarity: ClarityDetailed,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "unknown sub-guide",
			guides: map[string]Guide{
				"guide1": {
					ID: "guide1",
					Metadata: GuideMetadata{
						Scope:     ScopeLesson,
						Clarity:   ClarityDetailed,
						SubGuides: []string{"unknown_guide"},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "invalid sub-guide scope",
			guides: map[string]Guide{
				"guide1": {
					ID: "guide1",
					Metadata: GuideMetadata{
						Scope:     ScopeExplanation,
						Clarity:   ClarityDetailed,
						SubGuides: []string{"guide2"},
					},
				},
				"guide2": {
					ID: "guide2",
					Metadata: GuideMetadata{
						Scope:   ScopeLesson, // larger scope than Explanation
						Clarity: ClarityDetailed,
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateGraph(tt.guides)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateGraph() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
