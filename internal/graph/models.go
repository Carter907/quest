package graph

type Scope string

const (
	ScopeDefinition  Scope = "definition"
	ScopeDescription Scope = "description"
	ScopeExplanation Scope = "explanation"
	ScopeLesson      Scope = "lesson"
)

// ScopeValue helps with comparing scopes hierarchically.
func ScopeValue(s Scope) int {
	switch s {
	case ScopeDefinition:
		return 1
	case ScopeDescription:
		return 2
	case ScopeExplanation:
		return 3
	case ScopeLesson:
		return 4
	default:
		return 0
	}
}

type Clarity string

const (
	ClarityStrict       Clarity = "strict"
	ClarityDetailed     Clarity = "detailed"
	ClarityIntroductory Clarity = "introductory"
	ClarityVague        Clarity = "vague"
)

type GuideMetadata struct {
	Prerequisites []string `yaml:"prerequisites"`
	SubGuides     []string `yaml:"sub_guides"`
	Clarity       Clarity  `yaml:"clarity"`
	Scope         Scope    `yaml:"scope"`
	Tags          []string `yaml:"tags"`
}

type Guide struct {
	ID       string
	Path     string
	Metadata GuideMetadata
}
