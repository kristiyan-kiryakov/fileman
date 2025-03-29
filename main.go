package main

import (
	"fileman/cmd"
	"fmt"
	"os"
)

func main() {
	cmdInstance := cmd.NewRootCmd()
	if err := cmdInstance.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
