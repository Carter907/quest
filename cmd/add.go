package cmd

import (
	"fmt"
	"os"
	"strings"

	"charm.land/huh/v2"
	"github.com/Carter907/quest-cli/internal/graph"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v3"
)

var (
	interactive      bool
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
		if interactive {
			prereqString := ""
			subguideString := ""
			tagsString := ""
			var fields []huh.Field

			if !cmd.Flags().Changed("scope") {
				fields = append(fields, huh.NewSelect[string]().
					Title("Choose your Scope").
					Options(
						huh.NewOption("Definition", "definition"),
						huh.NewOption("Description", "description"),
						huh.NewOption("Explanation", "explanation"),
						huh.NewOption("Lesson", "lesson"),
					).
					Value(&addScope))
			}

			if !cmd.Flags().Changed("clarity") {
				fields = append(fields, huh.NewSelect[string]().
					Title("Choose a clarity").
					Options(
						huh.NewOption("Vague", "vague"),
						huh.NewOption("Introductory", "introductory"),
						huh.NewOption("Detailed", "detailed"),
						huh.NewOption("Strict", "strict"),
					).
					Value(&addClarity))
			}

			if !cmd.Flags().Changed("prerequisites") {
				fields = append(fields, huh.NewInput().
					Title("Enter prerequisites (separated by a comma)").
					Placeholder("e.g. Guide 1, Guide 2, Guide 3").
					Value(&prereqString))
			}

			if !cmd.Flags().Changed("subguides") {
				fields = append(fields, huh.NewInput().
					Title("Enter subguides (separated by a comma)").
					Placeholder("e.g. Guide 1, Guide 2, Guide 3").
					Value(&subguideString))
			}

			if !cmd.Flags().Changed("tags") {
				fields = append(fields, huh.NewInput().
					Title("Enter tags (comma separated)").
					Placeholder("e.g. Math, Science, Art").
					Value(&tagsString))
			}

			if len(fields) > 0 {
				form := huh.NewForm(
					huh.NewGroup(fields...),
				)

				err := form.Run()
				if err != nil {
					fmt.Printf("Failed to run form: %v\n", err)
					os.Exit(1)
				}
			}

			parseCSV := func(s string) []string {
				s = strings.TrimSpace(s)
				if s == "" {
					return []string{}
				}
				parts := strings.Split(s, ",")
				var res []string
				for _, p := range parts {
					if trimmed := strings.TrimSpace(p); trimmed != "" {
						res = append(res, trimmed)
					}
				}
				return res
			}

			if !cmd.Flags().Changed("prerequisites") {
				addPrerequisites = parseCSV(prereqString)
			}
			if !cmd.Flags().Changed("subguides") {
				addSubGuides = parseCSV(subguideString)
			}
			if !cmd.Flags().Changed("tags") {
				addTags = parseCSV(tagsString)
			}
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
	addCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "Interactive mode")
	addCmd.Flags().StringSliceVar(&addPrerequisites, "prerequisites", nil, "List of prerequisites")
	addCmd.Flags().StringSliceVar(&addSubGuides, "subguides", nil, "List of subguides")
	addCmd.Flags().StringVar(&addScope, "scope", "", "Scope of the guide (e.g. definition, description)")
	addCmd.Flags().StringVar(&addClarity, "clarity", "", "Clarity of the guide (e.g. strict, vague)")
	addCmd.Flags().StringSliceVar(&addTags, "tags", nil, "List of tags")

	rootCmd.AddCommand(addCmd)
}
