package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// freeCmd represents the free command
var freeCmd = &cobra.Command{
	Use:   "free",
	Short: "Unzip a .kng file so you can read or edit the knowledge graph",
	Long: `Free is how you start learning new knwoledge. .kng files get unzipped through free and 
	becomes accessible to the user.
	Examples:
		qst free my-graph.kng
	
	Unzipped archives are placed in a directory with the same name as the file. All guides and
	metadata losslessly decompressed.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("free called")
	},
}

func init() {
	rootCmd.AddCommand(freeCmd)
}
