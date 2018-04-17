package parser

import (
	"strings"
)

// GetCLICommands returns a slice of Commands
// it expects package strings to be formatted
// {name}={user}/{image}:{tag}, where image is
// the only required field
// TODO: add validation
func GetCLICommands(packages []string) map[string]Command {
	commands := map[string]Command{}

	for _, pkg := range packages {
		command := getCLICommand(pkg)
		commands[command.Name] = command
	}

	return commands
}

func getCLICommand(pkg string) Command {
	command := Command{}

	name, image := getPackageNameAndImage(pkg)
	command.Name = name
	command.Image = image

	return command
}

// if no given name, find name from image
// else use assigned name
func getPackageNameAndImage(pkg string) (string, string) {
	var name, image string
	strs := strings.Split(pkg, "=")

	if len(strs) == 1 {
		name = getNameFromImage(pkg)
		image = pkg
	} else {
		name = strs[0]
		image = strs[1]
	}

	image = addTagIfNotGiven(image)

	return name, image
}

func getNameFromImage(pkg string) string {
	strs := strings.Split(pkg, "/")

	if len(strs) != 1 {
		pkg = strs[1]
	}

	strs = strings.Split(pkg, ":")
	return strs[0]
}

func addTagIfNotGiven(image string) string {
	i := strings.Index(image, ":")

	if i == -1 {
		image = image + ":latest"
	}

	return image
}
