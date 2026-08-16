package cmd

import (
	"fmt"
	"os"

	"github.com/Carter907/quest-cli/internal/graph"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v3"
)

var (
	addPrerequisites []string
	addSubGuides     []string
	addScope         string
	addClarity       string
	addTags          []string
)

var addCmd = &cobra.Command{
	Use:   "add [name]",
	Short: "Add a new guide to the knowledge graph.",
	Long:  "Add allows you to insert a guide into the knowledge graph by specifying it's prerequisites, subguides, scope, and clarity. A new markdown file will be inserted into the directory",
	Example: `# Add a new definition
# missing value flags correspond to empty properties
qst add Exponent --scope definition --clarity strict`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		filename := name + ".md"

		if _, err := os.Stat(filename); !os.IsNotExist(err) {
			fmt.Printf("Error: %s already exists\n", filename)
			os.Exit(1)
		}

		meta := graph.GuideMetadata{
			Prerequisites: addPrerequisites,
			SubGuides:     addSubGuides,
			Scope:         graph.Scope(addScope),
			Clarity:       graph.Clarity(addClarity),
			Tags:          addTags,
		}

		// Ensure nil slices serialize to empty arrays [] in yaml instead of null
		if meta.Prerequisites == nil {
			meta.Prerequisites = []string{}
		}
		if meta.SubGuides == nil {
			meta.SubGuides = []string{}
		}
		if meta.Tags == nil {
			meta.Tags = []string{}
		}

		metaBytes, err := yaml.Marshal(&meta)
		if err != nil {
			fmt.Printf("Failed to marshal frontmatter: %v\n", err)
			os.Exit(1)
		}

		content := fmt.Sprintf("---\n%s---\n\n## %s\n\nThis is a guide for %s.\n", string(metaBytes), name, name)

		err = os.WriteFile(filename, []byte(content), 0o644)
		if err != nil {
			fmt.Printf("Failed to write %s: %v\n", filename, err)
			os.Exit(1)
		}

		fmt.Printf("Successfully added new guide: %s\n", filename)
	},
}

func init() {
	addCmd.Flags().StringSliceVar(&addPrerequisites, "prerequisites", nil, "List of prerequisites")
	addCmd.Flags().StringSliceVar(&addSubGuides, "subguides", nil, "List of subguides")
	addCmd.Flags().StringVar(&addScope, "scope", "", "Scope of the guide (e.g. definition, description)")
	addCmd.Flags().StringVar(&addClarity, "clarity", "", "Clarity of the guide (e.g. strict, vague)")
	addCmd.Flags().StringSliceVar(&addTags, "tags", nil, "List of tags")

	rootCmd.AddCommand(addCmd)
}
