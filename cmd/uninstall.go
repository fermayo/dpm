package cmd

import (
	"fmt"
	"os"

	"github.com/fermayo/dpm/project"
	"github.com/spf13/cobra"
)

var forceUninstall bool

func init() {
	uninstallCmd.Flags().BoolVarP(&forceUninstall, "force", "f", false,
		"Force uninstall even if project is currently active")
	RootCmd.AddCommand(uninstallCmd)
}

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstalls all commands for the current project",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: add option to remove images
		// will need to implement something that shows what
		// images are used by each project

		os.RemoveAll(project.ProjectCmdPath)
		fmt.Println("All commands uninstalled")
	},
}
