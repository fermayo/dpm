package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCLICommands(t *testing.T) {
	t.Parallel()

	inputOne := "node:6"
	outputOne := Command{
		Name:  "node",
		Image: "node:6",
		Entrypoints: []string{
			"node",
		},
	}

	inputTwo := "python"
	outputTwo := Command{
		Name:  "python",
		Image: "python:latest",
		Entrypoints: []string{
			"python",
		},
	}

	inputThree := "go=golang:1.7.5"
	outputThree := Command{
		Name:  "go",
		Image: "golang:1.7.5",
		Entrypoints: []string{
			"go",
		},
	}

	inputFour := "test=JPZ13/foo:13"
	outputFour := Command{
		Name:  "test",
		Image: "JPZ13/foo:13",
		Entrypoints: []string{
			"test",
		},
	}

	inputFive := "JPZ13/bar:13"
	outputFive := Command{
		Name:  "bar",
		Image: "JPZ13/bar:13",
		Entrypoints: []string{
			"bar",
		},
	}

	inputSlice := []string{
		inputOne,
		inputTwo,
		inputThree,
		inputFour,
		inputFive,
	}

	commandMap := GetCommandsFromCLI(inputSlice)
	require.Equal(t, outputOne, commandMap[outputOne.Name])
	require.Equal(t, outputTwo, commandMap[outputTwo.Name])
	require.Equal(t, outputThree, commandMap[outputThree.Name])
	require.Equal(t, outputFour, commandMap[outputFour.Name])
	require.Equal(t, outputFive, commandMap[outputFive.Name])
}
