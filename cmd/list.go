/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/luckuveryx/tix/internal/tix"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all Tixs",
	Aliases: []string{"l"},
	RunE:    runList,
}

func runList(cmd *cobra.Command, args []string) error {
	var tixs tix.Tixs
	if err := tixStorage.Load(&tixs); err != nil {
		fmt.Println("Error loading tixs")
		return err
	}
	tixs.Print()
	return nil

}

func init() {
	rootCmd.AddCommand(listCmd)
}
