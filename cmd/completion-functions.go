package cmd

import (
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

func completeModules(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	modules, err := repo.ListModules()
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}

	modules2 := []string{}
	for _, module := range modules {
		if strings.HasPrefix(module, toComplete) && slices.Index(args, module) == -1 {
			modules2 = append(modules2, module)
		}
	}

	if len(modules2) == 0 {
		return modules2, cobra.ShellCompDirectiveError
	}
	return modules2, cobra.ShellCompDirectiveDefault
}
