package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "info the dotfiles repository",
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		absoluteRepoLocation, err := filepath.Abs(repoLocation)
		cobra.CheckErr(err)

		fmt.Printf("\x1b[32mrepo\x1b[0m            : %s\n", repoLocation)
		fmt.Printf("\x1b[32mrepo (absolute)\x1b[0m : %s\n", absoluteRepoLocation)
		fmt.Printf("\x1b[32mtarget\x1b[0m          : %s\n", targetDir)

		for _, arg := range args {
			stat, err := os.Lstat(arg)
			if err != nil {
				fmt.Printf("Failed to stat %s\n", arg)
				continue
			}
			isSymlink := stat.Mode()&os.ModeSymlink != 0
			if !isSymlink {
				continue
			}

			dest, err := os.Readlink(arg)
			if !path.IsAbs(dest) {
				dir, err := filepath.Abs(path.Dir(arg))
				if err != nil {
					panic(err)
				}
				dest = path.Join(dir, path.Dir(dest))
			}
			fmt.Printf("%s: %v\n", arg, strings.HasPrefix(dest, absoluteRepoLocation))
		}

		return nil
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return nil, cobra.ShellCompDirectiveDefault
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
