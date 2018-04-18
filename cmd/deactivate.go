package cmd

import (
	"fmt"
	"log"

	"github.com/fermayo/dpm/alias"
	"github.com/fermayo/dpm/project"
	"github.com/fermayo/dpm/switcher"
	"github.com/spf13/cobra"
)

var forceDeactivate bool

func init() {
	deactivateCmd.Flags().BoolVarP(&forceDeactivate, "force", "f", false,
		"Force deactivation even if another project is currently active")
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

		if !forceDeactivate {
			if switchProjectName == "" {
				log.Fatal("error: no active project\n")
			}

			if switchProjectName != project.ProjectName {
				log.Fatalf("error: project '%s' already active", switchProjectName)
			}
		}

		err = switcher.UnsetSwitch()
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		err = alias.UnsetAliases()
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		if switchProjectName != "" {
			fmt.Printf("Project '%s' deactivated\n", switchProjectName)
		} else {
			fmt.Print("No active project to deactivate\n")
		}
	},
}
