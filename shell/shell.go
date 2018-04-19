package shell

import (
	"fmt"
	"os"
)

func StartShell() error {
	// Get the current working directory.
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	// Set an environment variable.
	os.Setenv("DPM_ACTIVE", "1")

	// Transfer stdin, stdout, and stderr to the new process
	// and also set target directory for the shell to start in.
	pa := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Dir:   cwd,
	}

	// Start up a new shell.
	fmt.Print(">>> Starting a new interactive shell")

	shell := os.Getenv("SHELL")
	proc, err := os.StartProcess(shell, []string{shell}, &pa)
	if err != nil {
		return err
	}

	// Wait until user exits the shell
	state, err := proc.Wait()
	if err != nil {
		return err
	}

	fmt.Printf("<<< Exited shell: %s\n", state.String())
	return nil
}
