package project

import (
	"io/ioutil"
	"path"

	"github.com/fermayo/dpm/utils"
)

func ActivateProject() error {
	configFileName, err := makeConfigIfNotExist()
	if err != nil {
		return err
	}

	return activateProjectInConfig(configFileName)
}

func IsProjectActive() (bool, error) {
	homeDir, err := getHomeDirectory()
	if err != nil {
		return false, err
	}

	filename := path.Join(homeDir, configName)
	doesExist, err := utils.DoesFileExist(filename)
	if err != nil {
		return false, err
	}

	if !doesExist {
		return false, nil
	}

	projectTable, err := getProjectTable(filename)
	if err != nil {
		return false, err
	}

	return projectTable[ProjectFilePath], nil
}

func activateProjectInConfig(filename string) error {
	projectTable, err := getProjectTable(filename)
	if err != nil {
		return err
	}

	projectTable[ProjectFilePath] = true

	return writeProjectTableToFile(projectTable, filename)
}

func makeConfigIfNotExist() (string, error) {
	homeDir, err := getHomeDirectory()
	if err != nil {
		return "", err
	}

	filename := path.Join(homeDir, configName)
	doesExist, err := utils.DoesFileExist(filename)
	if err != nil {
		return "", err
	}

	if !doesExist {
		return filename, ioutil.WriteFile(filename, []byte("{}"), 0755)
	}

	return filename, nil
}
