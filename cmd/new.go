package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "create a new knowledge graph directory",
	Long: `You start a new knowledge graph with the new command. You only have to specify the name of the knowledge graph;
	a directory will be created. A starter guide will be added in the directory as a template.

	If you don't specify a directory, the current directory will be used.

	Examples:

		qst new learn-cpp

	if you don't specify a directory, the current one will be used:

		qst new
	`,
	Run: func(cmd *cobra.Command, args []string) {
		dir := "."
		if len(args) != 0 {
			dir = args[0]
		}

		err := os.MkdirAll(dir, 0o755)
		if err != nil {
			fmt.Printf("Failed to create knowledge graph: %v\n", err)
			os.Exit(1)
		}

		starterPath := filepath.Join(dir, "starter.md")
		if _, err := os.Stat(starterPath); os.IsNotExist(err) {
			starterContent := `---
prerequisites: []
sub_guides: []
clarity: strict
scope: definition
tags: ["example"]
---

## Starter Guide

This is an example guide. Replace this content with your own knowledge!
`
			err = os.WriteFile(starterPath, []byte(starterContent), 0o644)
			if err != nil {
				fmt.Printf("Failed to create starter guide: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Initialized empty knowledge graph in %s\n", dir)
			fmt.Printf("Created starter guide at %s\n", starterPath)
		} else {
			fmt.Printf("Knowledge graph initialized in %s\n", dir)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
