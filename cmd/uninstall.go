package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func init() {
	RootCmd.AddCommand(uninstallCmd)
}

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstalls all commands for the current project",
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		cmdDir := fmt.Sprintf("%s/.dpm", dir)
		os.RemoveAll(cmdDir)
		fmt.Println("All commands uninstalled")
	},
}
