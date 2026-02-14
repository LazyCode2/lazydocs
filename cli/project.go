package cli

import (
	"time"

	"codeberg.org/LazyCode2/lazydocs/core"
	"codeberg.org/LazyCode2/lazydocs/site/template"
)

func InitProject(version string) error {
	docs := &core.APIDocumentation{
		Version:   version,
		Endpoints: []core.APIEndpoint{},
		UpdatedAt: time.Now(),
	}

	if err := core.SaveDocs(docs); err != nil {
		logger.Warn("Failed to initialize docs: %v", err)
		return err
	}

	if err := template.GenerateTemplate(); err != nil {
		logger.Warn("Failed to create template")
		return err
	}

	logger.Success("Initialized new lazydocs project")
	logger.Success("Template initialized")
	logger.Info("Docs file created")

	return nil
}

func AddBaseUrl(url string) error {
	docs, err := core.LoadDocs()
	if err != nil {
		logger.Warn("Failed to load docs: %v", err)
		return err
	}

	docs.BaseURL = url
	docs.UpdatedAt = time.Now()

	if err := core.SaveDocs(docs); err != nil {
		logger.Warn("Failed to save docs: %v", err)
		return err
	}

	logger.Success("Base URL updated successfully")
	return nil
}
