package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCommands(t *testing.T) {
	commandMap := GetCommands("./test-data/dpm.yml")
	glideCommand := Command{
		Name:    "glide",
		Image:   "dockerepo/glide",
		Context: "/run/context",
		Entrypoints: []string{
			"glide",
		},
	}

	goCommand := Command{
		Name:    "go",
		Image:   "golang:1.7.5",
		Context: "/go/src/github.com/JPZ13/dpm",
		Entrypoints: []string{
			"go",
		},
	}

	pythonCommand := Command{
		Name:    "python",
		Image:   "python:latest",
		Context: "/run/context",
		Entrypoints: []string{
			"python",
		},
	}

	require.Equal(t, glideCommand, commandMap[glideCommand.Name])
	require.Equal(t, goCommand, commandMap[goCommand.Name])
	require.Equal(t, pythonCommand, commandMap[pythonCommand.Name])
}
