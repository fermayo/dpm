package alias

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/fermayo/dpm/parser"
	"github.com/fermayo/dpm/project"
	"github.com/fermayo/dpm/utils"
)

type aliasSetter bool

const (
	set   aliasSetter = true
	unset aliasSetter = false
)

const (
	bashFile       = "#!/bin/bash\nif [ \"$DPM_ACTIVE\" == '1' ]; then\nexec \"%s/%s\" \"$@\"\nelse\nexec %s/%s-home \"$@\"\nfi"
	binaryLocation = "/usr/local/bin" // TODO: make variable for Windows
)

// SetAliases aliases all of the commands
// to use containers
func SetAliases() error {
	commandMap := parser.GetCommands(project.ProjectFilePath)

	return setOrUnsetAliases(commandMap, set)
}

// UnsetAliases removes all of the aliases
// set by SetAliases
func UnsetAliases() error {
	commandMap := parser.GetCommands(project.ProjectFilePath)

	return setOrUnsetAliases(commandMap, unset)
}

// setOrUnsetAliases loops the commands
func setOrUnsetAliases(aliases map[string]parser.Command, setter aliasSetter) error {
	for alias, _ := range aliases {
		err := setOrUnsetAlias(alias, setter)
		if err != nil {
			return err
		}
	}

	return nil
}

// setOrUnsetAlias invokes setAlias or unsetAlias
// depending on value passed to setter
func setOrUnsetAlias(alias string, setter aliasSetter) error {
	if setter == set {
		return setAlias(alias)
	}

	return unsetAlias(alias)
}

// unsetAlias removes the new alias in /usr/local/bin
// and restores the old alias
func unsetAlias(alias string) error {
	file := fmt.Sprintf("%s-home", alias)
	filename := path.Join(binaryLocation, file)
	doesExist, err := utils.DoesFileExist(filename)
	if err != nil {
		return err
	}

	new := path.Join(binaryLocation, alias)
	if doesExist {
		old := fmt.Sprintf("%s/%s-home", binaryLocation, alias)
		return os.Rename(old, new)
	}

	return os.Remove(new)
}

// setAlias adds a new alias in /usr/local/bin
// and re-aliases previously existing command
func setAlias(alias string) error {
	err := moveExistingAlias(alias)
	if err != nil {
		return err
	}

	err = generateBashFile(alias)
	if err != nil {
		return err
	}

	return nil
}

// generateBashFile makes a bash file that
// maps the command to what's in the project
// .dpm folder
func generateBashFile(alias string) error {
	contents := fmt.Sprintf(bashFile, project.ProjectCmdPath, alias, binaryLocation, alias)

	targetPath := path.Join(binaryLocation, alias)
	err := ioutil.WriteFile(targetPath, []byte(contents), 0755)
	if err != nil {
		return err
	}

	return nil
}

// moveExistingAlias moves the old command
// to a new file
func moveExistingAlias(alias string) error {
	filename := path.Join(binaryLocation, alias)
	doesExist, err := utils.DoesFileExist(filename)
	if err != nil {
		return err
	}

	if !doesExist {
		return nil
	}

	old := path.Join(binaryLocation, alias)
	new := fmt.Sprintf("%s/%s-home", binaryLocation, alias)
	return os.Rename(old, new)
}