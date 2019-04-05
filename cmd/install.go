package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/JPZ13/dpm/internal/parser"
	"github.com/JPZ13/dpm/internal/project"
	"github.com/spf13/cobra"
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

		os.RemoveAll(project.ProjectCmdPath)
		err := os.MkdirAll(project.ProjectCmdPath, 0755)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		if len(args) > 0 {
			err = installListedPackages(args)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
		}

		err = installYAMLPackages()
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	},
}

// installYAMLPackages writes bash scripts for
// each of the commands listed in the dpm.yml file
func installYAMLPackages() error {
	commands := parser.GetCommands(project.ProjectFilePath)
	data, err := ioutil.ReadFile(project.ProjectFilePath)
	if err != nil {
		return err
	}

	// TODO: figure out what this is used for
	err = ioutil.WriteFile(path.Join(project.ProjectCmdPath, "dpm.yml"), data, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Installing %d commands...\n", len(commands))

	commandNames := []string{}

	for name, command := range commands {
		commandNames = append(commandNames, name)
		cliCommands := commandToDockerCLIs(command)
		err = writeDockerBashCommands(cliCommands)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Installed: %s\n", strings.Join(commandNames, ", "))

	fmt.Print("Now you can run `dpm activate` to start using your new commands\n")

	return nil
}

// installListedPackages adds packages listed after
// install on the CLI to the dpm.yml file
func installListedPackages(packages []string) error {
	commands := parser.GetCommandsFromCLI(packages)

	return parser.AddCommands(project.ProjectFilePath, commands)
}

// commandToDockerCLIs takes a command and translates it into
// a docker cli command. It returns a map of the entrypoint to the
// docker command
func commandToDockerCLIs(command parser.Command) map[string]string {
	volumes := ""
	for _, volume := range command.Volumes {
		volumes = fmt.Sprintf("%s -v %s", volumes, volume)
	}

	cliCommands := make(map[string]string)
	for _, entrypoint := range command.Entrypoints {
		cliCommands[entrypoint] = fmt.Sprintf("docker run -it --rm -v $(pwd):%s %s -w %s --entrypoint %s %s \"$@\"",
			command.Context, volumes, command.Context, entrypoint, command.Image)
	}

	return cliCommands
}

// writeDockerBashCommands :
func writeDockerBashCommands(cliCommands map[string]string) error {
	for entrypoint, bashCommand := range cliCommands {
		targetPath := path.Join(project.ProjectCmdPath, entrypoint)
		contents := fmt.Sprintf("#!/bin/sh\nexec %s", bashCommand)

		err := ioutil.WriteFile(targetPath, []byte(contents), 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
