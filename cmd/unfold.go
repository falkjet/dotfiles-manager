package cmd

import (
	"github.com/spf13/cobra"
)

var unfoldCmd = &cobra.Command{
	Use:   "unfold",
	Short: "unfold one or more symlinks",
	RunE: func(cmd *cobra.Command, args []string) error {
		cobra.CheckErr(repo.LoadConfig(true))
		for _, arg := range args {
			err := repo.Unfold(arg)
			if err != nil {
				cmd.PrintErrln(err)
			}
		}
		return nil
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return nil, cobra.ShellCompDirectiveDefault
	},
}

func init() {
	rootCmd.AddCommand(unfoldCmd)
}
