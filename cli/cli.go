package cli

import (
	"flag"
)

var (
	addEndpoint = flag.String("add", "", "API endpoint to add (e.g., api/v1/homepage)")
)

func InitCli() {
	flag.Parse()
}

func ExecuteCommand() error {
	if *addEndpoint != "" {
		return handleAdd()
	}

	//Show usage if addEndpoint is ""
	flag.Usage()
	return nil
}

func handleAdd() error {
	//TODO Handle add command

	return nil
}
