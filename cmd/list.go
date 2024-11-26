package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all modules in dotfiles repository.",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		modules, err := repo.ListModules()
		if err != nil {
			return err
		}
		for _, module := range modules {
			fmt.Println(module)
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
