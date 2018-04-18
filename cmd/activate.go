package cmd

import (
	"fmt"
	"log"

	"github.com/fermayo/dpm/alias"
	"github.com/fermayo/dpm/project"
	"github.com/fermayo/dpm/switcher"
	"github.com/spf13/cobra"
)

var forceActivate bool

func init() {
	activateCmd.Flags().BoolVarP(&forceActivate, "force", "f", false,
		"Force activation even if another project is currently active")
	RootCmd.AddCommand(activateCmd)
}

var activateCmd = &cobra.Command{
	Use:   "activate",
	Short: "Activates the project in the current shell",
	Run: func(cmd *cobra.Command, args []string) {
		if !project.IsProjectInstalled() {
			// TODO: Just install the project instead of kicking error
			log.Fatal("error: commands are not installed - please run `dpm install` first\n")
		}

		switchProjectName, err := switcher.GetSwitchProjectName()
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		if !forceActivate && switchProjectName != "" {
			if switchProjectName == project.ProjectName {
				log.Fatal("error: current project already active\n")
			} else {
				log.Fatalf("error: project '%s' already active", switchProjectName)
			}
		}

		err = switcher.SetSwitch(project.ProjectCmdPath)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		err = alias.SetAliases()
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		fmt.Printf("Project '%s' activated\n", project.ProjectName)
	},
}
