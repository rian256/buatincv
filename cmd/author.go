package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(authorCmd)
}

var author = "Rian"

var authorCmd = &cobra.Command{
	Use:   "author",
	Short: "Print the author of BUATINCV",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Author: %s\n", author)
	},
}
