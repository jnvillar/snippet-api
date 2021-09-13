package main

import (
	"os"

	"snippetapp/cmd"
)

func main() {
	if err := cmd.Cmds().Execute(); err != nil {
		os.Exit(1)
	}
}
