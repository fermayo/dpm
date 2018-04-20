package project

import (
	"errors"
	"path"

	"github.com/fermayo/dpm/utils"
)

const (
	projectNotActiveError = "Project already not active"
)

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

	projectTable[ProjectFilePath] = false

	return writeProjectTableToFile(projectTable, filename)
}