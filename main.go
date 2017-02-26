package main

import (
	"fmt"
	"github.com/fermayo/dpm/cmd"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
