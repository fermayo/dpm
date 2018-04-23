package alias

import (
	"fmt"
	"os"
	"path"

	"github.com/fermayo/dpm/parser"
	"github.com/fermayo/dpm/project"
	"github.com/fermayo/dpm/utils"
)

// UnsetAliases removes all of the aliases
// set by SetAliases
func UnsetAliases() error {
	commandMap := parser.GetCommands(project.ProjectFilePath)

	return setOrUnsetAliases(commandMap, unset)
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
