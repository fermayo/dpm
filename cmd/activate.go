package cmd

import (
	"fmt"
	"log"

	"github.com/JPZ13/dpm/internal/alias"
	"github.com/JPZ13/dpm/internal/project"
	"github.com/JPZ13/dpm/internal/shell"
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
			installYAMLPackages()
		}

		err := project.ActivateProject()
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		err = alias.SetAliases()
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		err = shell.StartShell(shell.Activate)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		fmt.Printf("Project '%s' activated\n", project.ProjectName)
	},
}
