package parser

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
)

var filename = "dpm.yml"

func GetCommands() map[string]Command {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	inputfile := map[string]map[string]Command{}
	err = yaml.Unmarshal([]byte(fileBytes), &inputfile)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	commands, ok := inputfile["commands"]
	if !ok {
		log.Fatal("no commands found in input file")
	}

	for name, command := range commands {
		command.Name = name

		if command.Context == "" {
			command.Context = "/run/context"
		}

		if command.Entrypoint == "" {
			command.Entrypoint = command.Name
		}

		commands[name] = command
	}

	return commands
}
