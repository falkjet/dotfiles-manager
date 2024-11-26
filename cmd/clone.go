package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone the dotfiles repository",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		err, exists := fileExists(repoLocation)
		cobra.CheckErr(err)
		if exists {
			return fmt.Errorf("%q already exists", repoLocation)
		}
		c := exec.Command("git", "clone", "--", args[0], repoLocation)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		cobra.CheckErr(c.Run())

		return nil
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}

func fileExists(filename string) (err error, exists bool) {
	_, err = os.Stat(filename)
	if os.IsNotExist(err) {
		return nil, false
	}
	if err != nil {
		return err, false
	}
	return nil, true
}
