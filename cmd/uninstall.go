package cmd

import (
	"fmt"
	"github.com/fermayo/dpm/project"
	"github.com/fermayo/dpm/switcher"
	"github.com/spf13/cobra"
	"log"
	"os"
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
		if !project.IsProjectInstalled() {
			log.Fatal("error: commands are not installed - please run `dpm install` first\n")
		}

		currentActiveProject, err := switcher.GetSwitchProjectName()
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		if !forceUninstall && currentActiveProject == project.ProjectName {
			log.Fatal("error: project is currently active - please run `dpm deactivate` first\n")
		}

		os.RemoveAll(project.ProjectCmdPath)
		fmt.Println("All commands uninstalled")
	},
}
