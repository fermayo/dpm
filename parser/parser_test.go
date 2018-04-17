package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCommands(t *testing.T) {
	commandMap := GetCommands("./test-data/dpm.yml")
	fmt.Println("SHOW COMMANDS", commandMap)
	glideCommand := Command{
		Name:       "glide",
		Image:      "dockerepo/glide",
		Context:    "/run/context",
		Entrypoint: "glide",
	}

	goCommand := Command{
		Name:       "go",
		Image:      "golang:1.7.5",
		Context:    "/go/src/github.com/fermayo/dpm",
		Entrypoint: "go",
	}

	pythonCommand := Command{
		Name:       "python",
		Image:      "python:latest",
		Context:    "/run/context",
		Entrypoint: "python",
	}

	require.Equal(t, glideCommand, commandMap[glideCommand.Name])
	require.Equal(t, goCommand, commandMap[goCommand.Name])
	require.Equal(t, pythonCommand, commandMap[pythonCommand.Name])
}
