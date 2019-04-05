package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JPZ13/dpm/cmd"
)

func main() {
	log.SetFlags(0)
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
