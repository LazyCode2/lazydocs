package cli

import (
	"fmt"
	"time"

	"codeberg.org/LazyCode2/lazydocs/core"
	"codeberg.org/LazyCode2/lazydocs/utils"
)

var logger = utils.NewLogger()

func AddAPIEndpoint(path, method, description, requestBody, responseBody string, statusCode int) error {
	docs, err := core.LoadDocs()
	if err != nil {
		return fmt.Errorf("failed to load docs: %w", err)
	}

	for i, ep := range docs.Endpoints {
		if ep.Path == path && ep.Method == method {
			docs.Endpoints[i].Description = description
			docs.Endpoints[i].RequestBody = requestBody
			docs.Endpoints[i].ResponseBody = responseBody
			docs.Endpoints[i].StatusCode = statusCode
			docs.Endpoints[i].UpdatedAt = time.Now()

			if err := core.SaveDocs(docs); err != nil {
				return fmt.Errorf("failed to save docs: %w", err)
			}

			fmt.Printf("✓ Updated endpoint: %s %s\n", method, path)
			return nil
		}
	}

	newEndpoint := core.APIEndpoint{
		Path:         path,
		Method:       method,
		Description:  description,
		RequestBody:  requestBody,
		ResponseBody: responseBody,
		StatusCode:   statusCode,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	docs.Endpoints = append(docs.Endpoints, newEndpoint)
	docs.UpdatedAt = time.Now()

	if err := core.SaveDocs(docs); err != nil {
		return fmt.Errorf("failed to save docs: %w", err)
	}

	logger.Info("Added endpoint: %s %s", method, path)
	logger.Info("Description: %s", description)
	logger.Info("Status: %d", statusCode)
	logger.Info("Docs saved to %s", "api-docs.json")

	return nil
}
