package cli

import (
	"flag"
	"fmt"
	"os"

	"codeberg.org/LazyCode2/lazydocs/core"
	"codeberg.org/LazyCode2/lazydocs/site"
)

var (
	addEndpoint  = flag.String("add", "", "API endpoint to add (e.g., api/v1/homepage)")
	method       = flag.String("m", "GET", "HTTP method (GET, POST, PUT, DELETE, PATCH)")
	description  = flag.String("d", "", "Endpoint description")
	tag          = flag.String("tag", " ", "Add tag to the API endpoint")
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
	listAll     = flag.Bool("list", false, "List all endpoints")
	listBy      = flag.String("lm", "", "List by method (GET, POST...)")
	showHelp    = flag.Bool("h", false, "Show usage")
	BaseUrl     = flag.String("base", "", "Add base url")

	// Site flags
	renderFlag = flag.Bool("render", false, "Render your API docs")
)

func InitCli() {
	flag.Parse()
}

func ExecuteCommand() error {
	switch {
	case *initProject:
		return InitProject(*versionFlag)

	case *listAll:
		return core.ViewAll()

	case *listBy != "":
		return core.ViewByMethod(*listBy)
	case *BaseUrl != "":
		return AddBaseUrl(*BaseUrl)

	case *renderFlag:
		if err := site.RenderStaticSite(); err != nil {
			fmt.Println("Failed to render site:", err)
			return err
		}
		logger.Info("Static site generated in /docs folder")
		return nil

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

	return AddAPIEndpoint(*addEndpoint, *method, *description, *tag, *requestBody, *responseBody, *statusCode)
}

func HandleDelete() error {
	return DeleteAPIEndpoint(*deleteEndpoint, *deleteMethod)
}

func HandleDeleteAll() error {
	logger.Warn("Are you sure you want to delete ALL endpoints? (yes/no): ")
	var confirmation string
	fmt.Scanln(&confirmation)

	if confirmation != "yes" {
		logger.Info("Deletion cancelled")
		return nil
	}

	return DeleteAllEndpoints()
}

func PrintHelp() {
	fmt.Println(`
lazydocs — Minimal API Docs CLI

Usage:
  lazydocs [flags]

Project:
  -init                    Initialize project
  -v <version>             API version (default: 1.0.0)
  -base <base_url>         Set base URL

View:
  -list                    List all API endpoints
  -lm <method>             List endpoints by HTTP method

Site:
	-render                  Generate documentation site from JSON

Add Endpoint:
  -add <endpoint>          Add new API endpoint
  -m <method>              HTTP method (default: GET)
  -d <description>         Endpoint description
  -body <schema>           Request body (optional)
  -response <schema>       Response body (optional)
  -status <code>           HTTP status code (default: 200)

Delete:
  -delete <endpoint>       Delete specific endpoint
  -dm <method>             Method of endpoint to delete
  -delete-all              Delete all endpoints

Other:
  -h                       Show help
`)
	os.Exit(0)
}
