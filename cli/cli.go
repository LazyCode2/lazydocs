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

	// Delete command flags
	deleteEndpoint = flag.String("delete", "", "API endpoint to delete (e.g., api/v1/homepage)")
	deleteMethod   = flag.String("dm", "GET", "HTTP method of endpoint to delete")
	deleteAll      = flag.Bool("delete-all", false, "Delete all endpoints")
)

func InitCli() {
	flag.Parse()
}

func ExecuteCommand() error {
	if *addEndpoint != "" {
		return handleAdd()
	}

	if *deleteAll {
		return HandleDeleteAll()
	}

	if *deleteEndpoint != "" {
		return HandleDelete()
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

func HandleDelete() error {
	return DeleteAPIEndpoint(*deleteEndpoint, *deleteMethod)
}

func HandleDeleteAll() error {
	fmt.Print("⚠️  Are you sure you want to delete ALL endpoints? (yes/no): ")
	var confirmation string
	fmt.Scanln(&confirmation)

	if confirmation != "yes" {
		fmt.Println("❌ Deletion cancelled")
		return nil
	}

	return DeleteAllEndpoints()
}
