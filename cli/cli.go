package cli

import (
	"flag"
	"fmt"
	"os"

	"codeberg.org/LazyCode2/lazydocs/core"
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

	// Project specific flags
	initProject = flag.Bool("init", false, "Initilize Project")
	versionFlag = flag.String("v", "1.0.0", "API version")
	listFlag    = flag.String("list", "", "List API endpoints. Optional method (GET, POST, etc.)")
	showHelp    = flag.Bool("h", false, "Show usage")
	BaseUrl     = flag.String("base", "", "Add base url")
)

func InitCli() {
	flag.Parse()
}

func ExecuteCommand() error {
	switch {
	case *initProject:
		return InitProject(*versionFlag)

	case *listFlag != "":
		if *listFlag == "" {
			return core.ViewAll()
		} else {
			return core.ViewByMethod(*listFlag)
		}

	case *BaseUrl != "":
		return AddBaseUrl(*BaseUrl)

	case *addEndpoint != "":
		return HandleAdd()

	case *deleteAll:
		return HandleDeleteAll()

	case *deleteEndpoint != "":
		return HandleDelete()

	case *showHelp:
		PrintHelp()
		return nil

	default:
		return nil
	}

}

func HandleAdd() error {
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

func PrintHelp() {
	fmt.Println(`
lazydocs — Minimal API Docs CLI

Usage:
  lazydocs [flags]

Flags:
  -init             Initialize project
  -v <version>      API version (default 1.0.0)
  -h				Show help

  -base <base_url>  Add base url for your endpoint

  -list 			List all API endpoints

  -add <endpoint>   Add new API endpoint
  -m <method>       HTTP method for add (default GET)
  -d <desc>         Description for add
  -body <schema>    Request body (optional)
  -response <schema> Response body (optional)
  -status <code>    HTTP status code (default 200)
  -delete <endpoint> Delete endpoint
  -dm <method>       Method of endpoint to delete
  -delete-all       Delete all endpoints
`)
	os.Exit(0)
}
