package project

import (
	"io/ioutil"
	"path"

	"github.com/JPZ13/dpm/internal/utils"
)

// ActivateProject makes a config file at
// $HOME/.dpm-config.json if it doesn't exist
// or sets the project path to true in the
// corresponding json
func ActivateProject() error {
	configFileName, err := makeConfigIfNotExist()
	if err != nil {
		return err
	}

	return activateProjectInConfig(configFileName)
}

// IsProjectActive reads the config  file
// at $HOME/.dpm-config.json to determine
// if the project is active
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

// activateProjectInConfig unmarshals the config json,
// sets the key of the project's file path the true,
// marshals the json, and writes the file like a gangsta
func activateProjectInConfig(filename string) error {
	projectTable, err := getProjectTable(filename)
	if err != nil {
		return err
	}

	projectTable[ProjectFilePath] = true

	return writeProjectTableToFile(projectTable, filename)
}

// makeConfigIfNotExist writes the config
// if the file does not exist
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
