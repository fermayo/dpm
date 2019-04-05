package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/JPZ13/dpm/internal/parser"
	"github.com/JPZ13/dpm/internal/project"
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
		if len(args) == 0 {
			uninstallAll()
		} else {
			err := uninstallListedPackages(args)
			if err != nil {
				log.Fatalf("error: %v", err)
			}

			installYAMLPackages()
		}
	},
}

func uninstallAll() {
	os.RemoveAll(project.ProjectCmdPath)
	fmt.Println("All commands uninstalled")
}

func uninstallListedPackages(packages []string) error {
	commands := parser.GetCommands(project.ProjectFilePath)

	for _, pkg := range packages {
		if _, ok := commands[pkg]; ok {
			delete(commands, pkg)
			continue
		}
		return errors.New(fmt.Sprintf("Command %s not in project", pkg))
	}

	return parser.UpdateCommands(project.ProjectFilePath, commands)
}
