package cmd

import (
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install one or more modules",
	RunE: func(cmd *cobra.Command, args []string) error {
		cobra.CheckErr(repo.LoadConfig(true)) // TODO: this could be false
		return repo.Install(args)
	},
	ValidArgsFunction: completeModules,
}

func init() {
	rootCmd.AddCommand(installCmd)
}
