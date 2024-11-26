package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "uninstall one or more modules",
	RunE: func(cmd *cobra.Command, args []string) error {
		return repo.Uninstall(args)
	},
	ValidArgsFunction: completeModules,
}

func uninstall(source, target string) error {

	sourceInfo, err := os.Stat(source)
	if err != nil {
		return err
	}
	targetInfo, err := os.Lstat(target)
	if err != nil {
		if os.IsNotExist(err) {
			os.Symlink(source, target)
			return nil
		} else {
			return err
		}
	}

	if sourceInfo.IsDir() {
		if !targetInfo.IsDir() {
			return fmt.Errorf("error")
		}

		entries, err := os.ReadDir(source)
		if err != nil {
			return err
		}
		for _, entry := range entries {
			s, t := path.Join(source, entry.Name()), path.Join(target, entry.Name())
			uninstall(s, t)
		}
	} else {
	}
	return nil
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
