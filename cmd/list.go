package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/JPZ13/dpm/internal/parser"
	"github.com/JPZ13/dpm/internal/project"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available commands in the current project",
	Run: func(cmd *cobra.Command, args []string) {
		isActive, err := project.IsProjectActive()
		if err != nil {
			log.Fatalf("error: %s", err)
		}

		if !isActive {
			log.Fatal("error: no active project - please run `dpm activate` first from your project root")
		}

		commands := parser.GetCommands(project.ProjectFilePath)
		w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
		fmt.Fprintln(w, "COMMAND\tIMAGE\tENTRYPOINT")
		for name, command := range commands {
			fmt.Fprintf(w, "%s\t%s\t%s\n", name, command.Image, command.Entrypoints)
		}
		w.Flush()
	},
}
