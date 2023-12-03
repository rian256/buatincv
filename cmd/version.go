package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionNumber = "1.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of BUATINCV",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("BUATINCV %s\n", versionNumber)
	},
}
