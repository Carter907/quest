/*
Package cmd

specifies base commands and flags for quest cli
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "qst",
	Short: "A lossless compression and formatting tool for knowledge",
	Long:  `qst can be used to initialize knowledge graph directories, update unformatted markdown files, and zip/unzip .kng files.`,
	// qst currently does nothing when invoked on it's own.
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Root returns the root command
func Root() *cobra.Command {
	return rootCmd
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.quest.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
