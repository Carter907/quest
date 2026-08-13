package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "create a new knowledge graph directory",
	Long: `this command is how you start a new knowledge graph from the command line. You specifiy the name
	of the knowledge graph and a directory will be created with a single guide using default values for
	the frontmatter. These default values can be configured. If you don't specify a directory, the current directory
	will be used.
	Examples:
		qst new learn-cpp

	if you don't specify a directory, the current one will be used:
		qst new

	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("new called")
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
