package cli

import (
	"flag"
	"fmt"
)

var (
	addEndpoint  = flag.String("add", "", "API endpoint to add (e.g., api/v1/homepage)")
	method       = flag.String("m", "GET", "HTTP method (GET, POST, PUT, DELETE, PATCH)")
	description  = flag.String("d", "", "Endpoint description")
	requestBody  = flag.String("body", "", "Request body schema (optional)")
	responseBody = flag.String("response", "", "Response body schema (optional)")
	statusCode   = flag.Int("status", 200, "Default status code")
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
	if *description == "" {
		return fmt.Errorf("description is required (use -d flag)")
	}

	return AddAPIEndpoint(*addEndpoint, *method, *description, *requestBody, *responseBody, *statusCode)
}
