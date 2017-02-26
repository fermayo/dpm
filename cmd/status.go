package cmd

import (
	"fmt"
	"github.com/fermayo/dpm/project"
	"github.com/fermayo/dpm/switcher"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	RootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Shows the current project status",
	Run: func(cmd *cobra.Command, args []string) {
		currentActiveProject, err := switcher.GetSwitchProjectName()
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		if currentActiveProject != "" {
			currentActiveProjectPath, err := switcher.GetSwitchProjectPath()
			if err != nil {
				log.Fatalf("error: %v", err)
			}

			currentActiveProject = fmt.Sprintf("%s (%s)", currentActiveProject, currentActiveProjectPath)
		} else {
			currentActiveProject = "(none)"
		}

		currentProject := "(no `dpm.yml` found in current directory)"
		if project.IsProjectInitialized() {
			currentProject = fmt.Sprintf("%s (%s)", project.ProjectName, project.ProjectPath)
		}

		fmt.Printf("Current project: %s\n", currentProject)
		fmt.Printf("Current active project: %s\n", currentActiveProject)
	},
}
