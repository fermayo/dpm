package parser

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/go-yaml/yaml"
)

func GetCommands(filename string) map[string]Command {
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
		log.Fatal("error: no commands found in input file")
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

func AddCommands(filename string, commands map[string]Command) error {
	inputFile, err := getInputFile(filename)
	if err != nil {
		return err
	}

	// append new commands
	oldCommands := inputFile["commands"]
	combinedMaps := appendCommandMaps(oldCommands, commands)

	// write to file
	inputFile["commands"] = combinedMaps
	return writeInputFile(filename, inputFile)
}

func UpdateCommands(filename string, commands map[string]Command) error {
	inputFile, err := getInputFile(filename)
	if err != nil {
		return err
	}

	inputFile["commands"] = commands
	return writeInputFile(filename, inputFile)
}

func getInputFile(filename string) (map[string]map[string]Command, error) {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	inputFile := map[string]map[string]Command{}
	err = yaml.Unmarshal([]byte(fileBytes), &inputFile)
	if err != nil {
		return nil, err
	}

	return inputFile, nil
}

func appendCommandMaps(mapOne, mapTwo map[string]Command) map[string]Command {
	combinedMaps := map[string]Command{}

	for name, command := range mapOne {
		combinedMaps[name] = command
	}

	for name, command := range mapTwo {
		combinedMaps[name] = command
	}

	return combinedMaps
}

func writeInputFile(filename string, inputFile map[string]map[string]Command) error {
	bytes, err := yaml.Marshal(inputFile)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, bytes, os.ModePerm)
}
