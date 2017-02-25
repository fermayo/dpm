package cmd

import (
	"fmt"
	"github.com/fermayo/dpm/parser"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func init() {
	RootCmd.AddCommand(installCmd)
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs all commands defined in dpm.yml in the current project",
	Run: func(cmd *cobra.Command, args []string) {
		commands := parser.GetCommands()
		dir, err := os.Getwd()
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		cmdDir := path.Join(dir, ".dpm")

		os.RemoveAll(cmdDir)
		err = os.MkdirAll(cmdDir, 0755)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		for name, command := range commands {
			fmt.Printf("Installing %s...\n", name)
			targetPath := path.Join(cmdDir, name)
			contents := fmt.Sprintf("#!/bin/sh\n%s", commandToDockerCLI(command))
			err = ioutil.WriteFile(targetPath, []byte(contents), 0755)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
		}
	},
}

func commandToDockerCLI(command parser.Command) string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return fmt.Sprintf("docker run -it --rm -v %s -w %s --entrypoint %s %s \"$@\"", fmt.Sprintf("%s:%s", dir, command.Context), command.Context, command.Entrypoint, command.Image)
}
