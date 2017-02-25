package cmd

import (
	"fmt"
	"github.com/fermayo/dpm/parser"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path"
	"text/tabwriter"
)

func init() {
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available commands in the current project",
	Run: func(cmd *cobra.Command, args []string) {
		commands := parser.GetCommands()
		dir, err := os.Getwd()
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		cmdDir := fmt.Sprintf("%s/.dpm", dir)

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
		fmt.Fprintln(w, "COMMAND\tIMAGE\tENTRYPOINT\tINSTALLED")
		for name, command := range commands {
			_, err := os.Stat(path.Join(cmdDir, name))
			installed := "Y"
			if err != nil {
				installed = "N"
			}

			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", name, command.Image, command.Entrypoint, installed)
		}
		w.Flush()
	},
}
