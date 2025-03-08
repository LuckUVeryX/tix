/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luckuveryx/tix/internal/tix"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:     "edit [index] [new title]",
	Short:   "Edit a Tix's title",
	Aliases: []string{"e"},
	RunE:    runEdit,
}

func runEdit(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("index and title are required")
	}

	idx, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	title := strings.Join(args[1:], " ")

	tixs := tix.Tixs{}
	if err := tixStorage.Load(&tixs); err != nil {
		return err
	}
	if err := tixs.Edit(idx, title); err != nil {
		return err
	}
	if err := tixStorage.Save(tixs); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(editCmd)
}
