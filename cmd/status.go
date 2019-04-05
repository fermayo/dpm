package cmd

import (
	"fmt"
	"log"

	"github.com/JPZ13/dpm/internal/project"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Shows the current project status",
	Run: func(cmd *cobra.Command, args []string) {
		isActive, err := project.IsProjectActive()
		if err != nil {
			log.Fatalf("error: %s", err)
		}

		msg := fmt.Sprintf("Project is active at %s", project.ProjectFilePath)
		if !isActive {
			msg = "No project active"
		}

		fmt.Println(msg)
	},
}
