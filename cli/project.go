package cli

import (
	"time"

	"codeberg.org/LazyCode2/lazydocs/core"
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

	logger.Success("Initialized new lazydocs project")
	logger.Info("Docs file created")

	return nil
}
