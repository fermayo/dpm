package project

import (
	"encoding/json"
	"io/ioutil"
	"os/user"
)

func writeProjectTableToFile(projectTable map[string]bool, filename string) error {
	tableBytes, err := json.Marshal(projectTable)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, tableBytes, 0755)
}

func getProjectTable(filename string) (map[string]bool, error) {
	configFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	projectTable := make(map[string]bool)
	err = json.Unmarshal(configFile, &projectTable)
	if err != nil {
		return nil, err
	}

	return projectTable, nil
}

func getHomeDirectory() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return usr.HomeDir, nil
}
