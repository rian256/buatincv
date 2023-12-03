package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "buatincv",
	Short: "Create your own resume inside CLI!",
	Long: `
			█▄▄ █░█ ▄▀█ ▀█▀ █ █▄░█ █▀▀ █░█
			█▄█ █▄█ █▀█ ░█░ █ █░▀█ █▄▄ ▀▄▀
        `,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func init() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
}
