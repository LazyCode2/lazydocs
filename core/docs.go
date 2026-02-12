package core

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

const (
	docsFileName = "api-docs.json"
)

type APIEndpoint struct {
	Path         string    `json:"path"`
	Method       string    `json:"method"`
	Description  string    `json:"description"`
	RequestBody  string    `json:"requestBody,omitempty"`
	ResponseBody string    `json:"responseBody,omitempty"`
	StatusCode   int       `json:"statusCode"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type APIDocumentation struct {
	Version   string        `json:"version"`
	Endpoints []APIEndpoint `json:"endpoints"`
	UpdatedAt time.Time     `json:"updatedAt"`
}

func LoadDocs() (*APIDocumentation, error) {
	if _, err := os.Stat(docsFileName); os.IsNotExist(err) {
		return &APIDocumentation{
			Version:   "1.0.0",
			Endpoints: []APIEndpoint{},
			UpdatedAt: time.Now(),
		}, nil
	}

	data, err := os.ReadFile(docsFileName)
	if err != nil {
		return nil, err
	}

	var docs APIDocumentation
	if err := json.Unmarshal(data, &docs); err != nil {
		return nil, err
	}

	return &docs, nil
}

func SaveDocs(docs *APIDocumentation) error {
	data, err := json.MarshalIndent(docs, "", "  ")
	if err != nil {
		return err
	}

	dir := filepath.Dir(docsFileName)
	if dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	if err := os.WriteFile(docsFileName, data, 0644); err != nil {
		return err
	}

	return nil
}
