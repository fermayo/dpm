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
	if err != nil {
		return "", err
	}

	projectPath, _ := path.Split(cmdPath)
	return path.Dir(projectPath), nil
}

func GetSwitchProjectName() (string, error) {
	_, err := os.Stat(SwitchPath)
	if err != nil {
		return "", nil
	}

	projectPath, _ := GetSwitchProjectPath()
	return path.Base(projectPath), nil
}

func SetSwitch(path string) error {
	return os.Symlink(path, SwitchPath)
}

func UnsetSwitch() error {
	return os.Remove(SwitchPath)
}
