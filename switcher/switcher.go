package switcher

import (
	"github.com/mitchellh/go-homedir"
	"log"
	"os"
	"path"
)

var SwitchPath string

func init() {
	// Set switcherPath
	homeDir, err := homedir.Dir()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	SwitchPath = path.Join(homeDir, ".dpm")
}

func GetSwitchProjectCmdPath() (string, error) {
	return os.Readlink(SwitchPath)
}

func GetSwitchProjectPath() (string, error) {
	cmdPath, err := GetSwitchProjectCmdPath()
	if os.IsNotExist(err) {
		return "", nil
	} else if err != nil {
		return "", err
	}

	projectPath, _ := path.Split(cmdPath)
	return path.Dir(projectPath), nil
}

func GetSwitchProjectName() (string, error) {
	_, err := os.Stat(SwitchPath)
	if os.IsNotExist(err) {
		return "", nil
	} else if err != nil {
		return "", err
	}

	projectPath, _ := GetSwitchProjectPath()
	return path.Base(projectPath), nil
}

func SetSwitch(path string) error {
	_, err := os.Stat(SwitchPath)
	if !os.IsNotExist(err) {
		return err
	} else {
		os.Remove(SwitchPath)
	}

	return os.Symlink(path, SwitchPath)
}

func UnsetSwitch() error {
	_, err := os.Stat(SwitchPath)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}

	return os.Remove(SwitchPath)
}
