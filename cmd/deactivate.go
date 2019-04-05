package cmd

import (
	"fmt"
	"log"

	"github.com/JPZ13/dpm/internal/alias"
	"github.com/JPZ13/dpm/internal/project"
	"github.com/JPZ13/dpm/internal/shell"
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
		err := project.DeactivateProject()
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		err = alias.UnsetAliases()
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		err = shell.StartShell(shell.Deactivate)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		fmt.Printf("Project '%s' deactivated\n", project.ProjectName)
	},
}
