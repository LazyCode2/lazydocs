package cli

import (
	"fmt"

	"codeberg.org/LazyCode2/lazydocs/core"
)

func DeleteAPIEndpoint(path, method string) error {
	docs, err := core.LoadDocs()
	if err != nil {
		return fmt.Errorf("failed to load docs: %w", err)
	}

	found := false
	for i, ep := range docs.Endpoints {
		if ep.Path == path && ep.Method == method {
			docs.Endpoints = append(docs.Endpoints[:i], docs.Endpoints[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("endpoint not found: %s %s", method, path)
	}

	if err := core.SaveDocs(docs); err != nil {
		return fmt.Errorf("failed to save docs: %w", err)
	}

	logger.Success("Deleted endpoint: %s %s", method, path)
	logger.Info("Docs saved to %s", "api-docs.json")

	return nil
}

func DeleteAllEndpoints() error {
	docs, err := core.LoadDocs()
	if err != nil {
		return fmt.Errorf("failed to load docs: %w", err)
	}

	if len(docs.Endpoints) == 0 {
		logger.Warn("No endpoints to delete")
		return nil
	}

	count := len(docs.Endpoints)
	docs.Endpoints = []core.APIEndpoint{}

	if err := core.SaveDocs(docs); err != nil {
		return fmt.Errorf("failed to save docs: %w", err)
	}

	logger.Success("Deleted %d endpoint(s)", count)
	logger.Info("Docs saved to %s", "api-docs.json")

	return nil
}
