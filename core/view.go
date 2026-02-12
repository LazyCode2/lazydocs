package core

import (
	"fmt"
)

func ViewAll() error {
	docs, err := LoadDocs()
	if err != nil {
		return fmt.Errorf("failed to load docs: %w", err)
	}

	for _, ep := range docs.Endpoints {
		fmt.Println("\033[36m---------------------------------------------------\033[0m")
		fmt.Printf("\033[33mPath:\033[0m %s\n", ep.Path)
		fmt.Printf("\033[32mMethod:\033[0m %s\n", ep.Method)
		fmt.Printf("\033[35mDescription:\033[0m %s\n", ep.Description)
		fmt.Printf("\033[31mStatus:\033[0m %d\n\n", ep.StatusCode)
	}

	return nil
}
