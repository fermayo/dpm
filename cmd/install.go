package cmd

import (
	"fmt"
	"github.com/fermayo/dpm/parser"
	"github.com/fermayo/dpm/project"
	"github.com/fermayo/dpm/switcher"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func init() {
	RootCmd.AddCommand(installCmd)
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs all commands defined in dpm.yml in the current project",
	Run: func(cmd *cobra.Command, args []string) {
		if !project.IsProjectInitialized() {
			log.Fatal("error: no `dpm.yml` file found\n")
		}

		commands := parser.GetCommands(project.ProjectFilePath)
		os.RemoveAll(project.ProjectCmdPath)
		err := os.MkdirAll(project.ProjectCmdPath, 0755)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		data, err := ioutil.ReadFile(project.ProjectFilePath)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		err = ioutil.WriteFile(path.Join(project.ProjectCmdPath, "dpm.yml"), data, 0644)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		fmt.Printf("Installing %d commands...\n", len(commands))
		commandNames := []string{}

		for name, command := range commands {
			commandNames = append(commandNames, name)
			targetPath := path.Join(project.ProjectCmdPath, name)
			contents := fmt.Sprintf("#!/bin/sh\nexec %s", commandToDockerCLI(command))
			err = ioutil.WriteFile(targetPath, []byte(contents), 0755)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
		}

		fmt.Printf("Installed: %s\n", strings.Join(commandNames, ", "))

		switchProjectName, err := switcher.GetSwitchProjectName()
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		if switchProjectName == "" {
			fmt.Print("Now you can run `dpm activate` to start using your new commands\n")
		}
	},
}

func commandToDockerCLI(command parser.Command) string {
	volumes := ""
	for _, volume := range command.Volumes {
		volumes = fmt.Sprintf("%s -v %s", volumes, volume)
	}
	return fmt.Sprintf("docker run -it --rm -v $(pwd):%s %s -w %s --entrypoint %s %s \"$@\"",
		command.Context, volumes, command.Context, command.Entrypoint, command.Image)
}
