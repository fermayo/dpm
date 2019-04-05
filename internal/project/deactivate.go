package project

import (
	"errors"
	"path"

	"github.com/JPZ13/dpm/internal/utils"
)

const (
	projectNotActiveError = "Project already not active"
)

// DeactivateProject removes the project from the config json file
func DeactivateProject() error {
	homeDir, err := getHomeDirectory()
	if err != nil {
		return err
	}

	filename := path.Join(homeDir, configName)

	doesExist, err := utils.DoesFileExist(filename)
	if !doesExist {
		return errors.New(projectNotActiveError)
	}

	projectTable, err := getProjectTable(filename)
	if err != nil {
		return err
	}

	if !projectTable[ProjectFilePath] {
		return errors.New(projectNotActiveError)
	}

	delete(projectTable, ProjectFilePath)

	return writeProjectTableToFile(projectTable, filename)
}
