package core

import "time"

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
