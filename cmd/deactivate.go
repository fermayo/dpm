package cmd

import (
	"fmt"
	"github.com/fermayo/dpm/project"
	"github.com/fermayo/dpm/switcher"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	RootCmd.AddCommand(deactivateCmd)
}

var deactivateCmd = &cobra.Command{
	Use:   "deactivate",
	Short: "Deactivates the project in the current shell",
	Run: func(cmd *cobra.Command, args []string) {
		switchProjectName, err := switcher.GetSwitchProjectName()
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		if switchProjectName == "" {
			log.Fatal("error: no project active\n")
		}

		if switchProjectName != project.ProjectName {
			log.Fatalf("error: project '%s' already active", switchProjectName)
		}

		err = switcher.UnsetSwitch()
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		fmt.Printf("Project '%s' deactivated\n", project.ProjectName)
	},
}
