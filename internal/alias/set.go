package alias

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/JPZ13/dpm/internal/parser"
	"github.com/JPZ13/dpm/internal/project"
	"github.com/JPZ13/dpm/internal/utils"
)

// SetAliases aliases all of the commands
// to use containers
func SetAliases() error {
	commandMap := parser.GetCommands(project.ProjectFilePath)

	return setOrUnsetAliases(commandMap, set)
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
	contents := fmt.Sprintf(bashFile, binaryLocation, alias, alias)

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
		return utils.WriteBashScript(filename, bashFileIfNotExist)
	}

	old := path.Join(binaryLocation, alias)
	new := fmt.Sprintf("%s/%s-home", binaryLocation, alias)
	return os.Rename(old, new)
}
