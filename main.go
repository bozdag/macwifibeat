package main

import (
	"os"

	"github.com/bozdag/macwifibeat/cmd"

	_ "github.com/bozdag/macwifibeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
