package cmd

import (
	"fmt"
	"strings"

	"github.com/luckuveryx/tix/internal/tix"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add [title]",
	Short:   "Add a new Tix",
	Aliases: []string{"a"},
	RunE:    runAdd,
}

func runAdd(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("title is required")
	}
	title := strings.Join(args, " ")
	tixs := &tix.Tixs{}
	if err := tixStorage.Load(tixs); err != nil {
		return err
	}
	tixs.Add(title)
	if err := tixStorage.Save(*tixs); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(addCmd)

}
