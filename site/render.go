package site

import (
	"html/template"
	"os"
	"path/filepath"
	"time"

	"codeberg.org/LazyCode2/lazydocs/core"
)

const outputDir = "docs"

type PageData struct {
	BaseURL   string
	Version   string
	UpdatedAt time.Time
	Endpoints []core.APIEndpoint
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

	tmpl := template.Must(template.New("index").Parse(htmlTemplate))

	filePath := filepath.Join(outputDir, "index.html")
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data := PageData{
		BaseURL:   docs.BaseURL,
		Version:   docs.Version,
		UpdatedAt: docs.UpdatedAt,
		Endpoints: docs.Endpoints,
	}

	return tmpl.Execute(file, data)
}

var htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>API Documentation</title>
	<style>
		body {
			font-family: Arial, sans-serif;
			margin: 40px;
			background: #f4f6f8;
		}
		h1 {
			color: #222;
		}
		.endpoint {
			background: white;
			padding: 20px;
			margin-bottom: 20px;
			border-radius: 8px;
			box-shadow: 0 2px 6px rgba(0,0,0,0.1);
		}
		.method {
			font-weight: bold;
			padding: 4px 8px;
			border-radius: 4px;
			color: white;
		}
		.GET { background: #16a34a; }
		.POST { background: #2563eb; }
		.PUT { background: #f59e0b; }
		.DELETE { background: #dc2626; }
		.PATCH { background: #7c3aed; }
		code {
			background: #eee;
			padding: 2px 6px;
			border-radius: 4px;
		}
	</style>
</head>
<body>

<h1>API Documentation</h1>
<p><strong>Base URL:</strong> {{.BaseURL}}</p>
<p><strong>Version:</strong> {{.Version}}</p>
<p><strong>Last Updated:</strong> {{.UpdatedAt}}</p>

<hr/>

{{range .Endpoints}}
<div class="endpoint">
	<div>
		<span class="method {{.Method}}">{{.Method}}</span>
		<code>{{.Path}}</code>
	</div>
	<p><strong>Description:</strong> {{.Description}}</p>

	{{if .RequestBody}}
	<p><strong>Request Body:</strong></p>
	<pre>{{.RequestBody}}</pre>
	{{end}}

	{{if .ResponseBody}}
	<p><strong>Response Body:</strong></p>
	<pre>{{.ResponseBody}}</pre>
	{{end}}

	<p><strong>Status Code:</strong> {{.StatusCode}}</p>
	<p><small>Created: {{.CreatedAt}} | Updated: {{.UpdatedAt}}</small></p>
</div>
{{end}}

</body>
</html>
`
