package site

import (
	"codeberg.org/LazyCode2/lazydocs/core"
	"html/template"
	"os"
	"path/filepath"
	"time"
)

const outputDir = "docs"

type PageData struct {
	BaseURL   string
	Version   string
	UpdatedAt string
	Endpoints []core.APIEndpoint
}

func LoadTemplate() (string, error) {
	data, err := os.ReadFile("template/index.tmpl")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func RenderStaticSite() error {
	docs, err := core.LoadDocs()
	if err != nil {
		return err
	}

	// Create output directory
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	tmplContent, err := LoadTemplate()
	if err != nil {
		return err
	}

	funcMap := template.FuncMap{
		"formatDate": func(t time.Time) string {
			return t.Format("Jan 2, 2006 at 3:04 PM")
		},
	}

	tmpl := template.Must(template.New("index").Funcs(funcMap).Parse(tmplContent))

	filePath := filepath.Join(outputDir, "index.html")
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data := PageData{
		BaseURL:   docs.BaseURL,
		Version:   docs.Version,
		UpdatedAt: docs.UpdatedAt.Format("Jan 2, 2006 at 3:04 PM"),
		Endpoints: docs.Endpoints,
	}

	return tmpl.Execute(file, data)
}
