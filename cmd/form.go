package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var formCmd = &cobra.Command{
	Use:   "form",
	Short: "zip a knowledge graph directory into the .kng archive file format",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("form called")
	},
}

func init() {
	rootCmd.AddCommand(formCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
