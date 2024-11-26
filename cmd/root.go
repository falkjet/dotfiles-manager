package cmd

import (
	"os"
	"path"

	"github.com/falkjet/dotfiles-manager/internal"
	"github.com/spf13/cobra"
)

var repoLocation string
var targetDir string
var repo *internal.Repo

var rootCmd = &cobra.Command{
	Use:   "dotfiles-manager",
	Short: "",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(
		&repoLocation, "repo-location", "",
		"The location of the dotfiles repo (default $HOME/dotfiles)")

	rootCmd.PersistentFlags().StringVar(
		&targetDir, "target", "",
		"The location of the dotfiles repo (default $HOME)")
}

func initConfig() {
	if repoLocation == "" {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		repoLocation = path.Join(home, "dotfiles")
	}
	if targetDir == "" {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		targetDir = home
	}

	repo = internal.NewRepo(repoLocation, targetDir)
}
