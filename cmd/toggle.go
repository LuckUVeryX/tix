/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/luckuveryx/tix/internal/tix"
	"github.com/spf13/cobra"
)

var toggleCmd = &cobra.Command{
	Use:     "toggle [index]",
	Short:   "Toggle completion status of a Tix",
	Aliases: []string{"t"},
	RunE:    runToggle,
}

func runToggle(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("index is required")
	}
	idx, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	tixs := tix.Tixs{}
	if err := tixStorage.Load(&tixs); err != nil {
		return err
	}
	if err := tixs.Toggle(idx); err != nil {
		return err
	}
	if err := tixStorage.Save(tixs); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(toggleCmd)
}
