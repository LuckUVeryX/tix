package cmd

import (
	"os"

	"github.com/luckuveryx/tix/internal/tix"
	"github.com/spf13/cobra"
)

var tixStorage *tix.Storage[tix.Tixs]

var rootCmd = &cobra.Command{
	Use:   "tix",
	Short: "Tix is a CLI todo app",
	Long: `Tix is a command-line todo application built in Go.
It allows you to manage your Tixs from the terminal.`,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	RunE: runList,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	tixStorage = tix.NewStorage[tix.Tixs]("/Users/ryanyip/Repositories/personal/tix/tixs.json")
}
