package cmd

import (
	"fmt"
	"github.com/fermayo/dpm/parser"
	"github.com/fermayo/dpm/switcher"
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
		switchProjectName, err := switcher.GetSwitchProjectName()
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		if switchProjectName == "" {
			log.Fatal("error: no active project - please run `dpm activate` first from your project root")
		}

		switchProjectPath, err := switcher.GetSwitchProjectPath()
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		commands := parser.GetCommands(path.Join(switchProjectPath, "dpm.yml"))
		w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
		fmt.Fprintln(w, "COMMAND\tIMAGE\tENTRYPOINT")
		for name, command := range commands {
			fmt.Fprintf(w, "%s\t%s\t%s\n", name, command.Image, command.Entrypoint)
		}
		w.Flush()
	},
}
