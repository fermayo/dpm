package project

import (
	"log"
	"os"
	"path"
)

var ProjectPath string
var ProjectCmdPath string
var ProjectFilePath string
var ProjectName string

func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	ProjectPath = wd
	ProjectCmdPath = path.Join(wd, ".dpm")
	ProjectFilePath = path.Join(wd, "dpm.yml")
	ProjectName = path.Base(ProjectPath)
}

func IsProjectInitialized() bool {
	_, err := os.Stat(ProjectFilePath)
	return err == nil
}

func IsProjectInstalled() bool {
	_, err := os.Stat(ProjectCmdPath)
	return err == nil
}
