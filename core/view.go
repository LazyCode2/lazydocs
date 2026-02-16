package core

import (
	"fmt"
	"os"

	"codeberg.org/LazyCode2/lazydocs/utils"
)

var logger = utils.NewLogger()

func ViewAll() error {
	docs, err := LoadDocs()
	if err != nil {
		return fmt.Errorf("failed to load docs: %w", err)
	}

	fmt.Printf("\033[36mBase Url:\033[0m %s\n", docs.BaseURL)
	for _, ep := range docs.Endpoints {
		fmt.Println("\033[36m---------------------------------------------------\033[0m")
		fmt.Printf("\033[33mPath:\033[0m %s\n", ep.Path)
		fmt.Printf("\033[32mMethod:\033[0m %s\n", ep.Method)
		fmt.Printf("\033[35mDescription:\033[0m %s\n", ep.Description)
		fmt.Printf("\033[31mStatus:\033[0m %d\n\n", ep.StatusCode)
	}

	return nil
}

func ViewByMethod(method string) error {
	docs, err := LoadDocs()
	if err != nil {
		return fmt.Errorf("failed to load docs: %w", err)
	}

	found := false // tracker for sort search

	fmt.Printf("\033[36mBase Url:\033[0m %s\n", docs.BaseURL)
	for _, ep := range docs.Endpoints {
		if ep.Method == method {
			found = true
			fmt.Println("\033[36m---------------------------------------------------\033[0m")
			fmt.Printf("\033[33mPath:\033[0m %s\n", ep.Path)
			fmt.Printf("\033[32mMethod:\033[0m %s\n", ep.Method)
			fmt.Printf("\033[35mDescription:\033[0m %s\n", ep.Description)
			fmt.Printf("\033[31mStatus:\033[0m %d\n\n", ep.StatusCode)
		}
	}

	if !found {
		logger.Warn("No endpoints found for method %s\n", method)
		os.Exit(1)
	}

	return nil
}
