package main

import (
	"fmt"
	"os"

	"codeberg.org/LazyCode2/lazydocs/cli"
)

func main() {
	cli.InitCli()

	if err := cli.ExecuteCommand(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
