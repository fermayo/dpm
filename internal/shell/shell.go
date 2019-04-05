package shell

import (
	"fmt"
	"os"

	"github.com/JPZ13/dpm/internal/project"
)

// ProjectState controls whether a project is active
type ProjectState bool

const (
	// Activate tells StartShell to activate env
	Activate ProjectState = true
	// Deactivate tells StartShell to unset env
	Deactivate ProjectState = false
)

// StartShell spawns a shell with updated
// environment variables
func StartShell(activate ProjectState) error {
	// Get the current working directory.
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	// Set an environment variable.
	if activate == Activate {
		os.Setenv("DPM_ACTIVE", project.ProjectCmdPath)
	} else if activate == Deactivate {
		os.Unsetenv("DPM_ACTIVE")
	}

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
