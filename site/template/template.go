package template

import "os"

func GenerateTemplate() error {
	if err := os.MkdirAll("template", 0755); err != nil {
		return err
	}

	err := os.WriteFile("template/index.tmpl", []byte(defaultTemplate), 0644)
	if err != nil {
		return err
	}

	return nil
}

const defaultTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>API Documentation - {{.Version}}</title>
	<style>
		* {
			margin: 0;
			padding: 0;
			box-sizing: border-box;
		}

		body {
			font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
			background: #ffffff;
			color: #37352f;
			line-height: 1.6;
			padding: 60px 20px;
		}

		.container {
			max-width: 900px;
			margin: 0 auto;
		}

		.header {
			margin-bottom: 48px;
		}

		.header h1 {
			font-size: 2.5rem;
			font-weight: 700;
			color: #37352f;
			margin-bottom: 8px;
		}

		.meta-info {
			display: flex;
			gap: 24px;
			flex-wrap: wrap;
			margin-top: 16px;
			color: #787774;
			font-size: 0.9rem;
		}

		.meta-item {
			display: flex;
			gap: 8px;
		}

		.meta-item strong {
			color: #37352f;
		}

		.divider {
			height: 1px;
			background: #e9e9e7;
			margin: 32px 0;
		}

		.endpoint {
			margin-bottom: 32px;
			padding-bottom: 32px;
			border-bottom: 1px solid #e9e9e7;
		}

		.endpoint:last-child {
			border-bottom: none;
		}

		.endpoint-header {
			display: flex;
			align-items: center;
			gap: 12px;
			margin-bottom: 12px;
		}

		.method {
			font-weight: 600;
			padding: 4px 10px;
			border-radius: 4px;
			font-size: 0.75rem;
			letter-spacing: 0.3px;
			text-transform: uppercase;
		}

		.GET { 
			background: #e9f3ec;
			color: #0f7b3f;
		}
		
		.POST { 
			background: #e7f3ff;
			color: #0b6bcb;
		}
		
		.PUT { 
			background: #fff4e5;
			color: #b35c00;
		}
		
		.DELETE { 
			background: #ffe9e9;
			color: #c41c1c;
		}
		
		.PATCH { 
			background: #f3f0ff;
			color: #6b46c1;
		}

		.path {
			font-family: 'SF Mono', Monaco, 'Cascadia Code', 'Courier New', monospace;
			color: #37352f;
			font-size: 0.95rem;
			background: #f7f6f3;
			padding: 4px 8px;
			border-radius: 4px;
		}

		.description {
			color: #37352f;
			font-size: 1rem;
			margin-bottom: 16px;
		}

		.section {
			margin-top: 16px;
		}

		.section-title {
			font-size: 0.85rem;
			font-weight: 600;
			color: #787774;
			margin-bottom: 8px;
			text-transform: uppercase;
			letter-spacing: 0.5px;
		}

		.code-block {
			background: #f7f6f3;
			color: #37352f;
			padding: 16px;
			border-radius: 4px;
			overflow-x: auto;
			font-family: 'SF Mono', Monaco, 'Cascadia Code', 'Courier New', monospace;
			font-size: 0.85rem;
			line-height: 1.5;
			border: 1px solid #e9e9e7;
			white-space: pre-wrap;
			word-wrap: break-word;
		}

		.status-code {
			display: inline-block;
			background: #f7f6f3;
			color: #37352f;
			padding: 4px 10px;
			border-radius: 4px;
			font-weight: 500;
			font-size: 0.9rem;
			font-family: 'SF Mono', Monaco, monospace;
			border: 1px solid #e9e9e7;
		}

		.timestamp {
			margin-top: 16px;
			color: #9b9a97;
			font-size: 0.8rem;
			display: flex;
			gap: 16px;
		}

		@media (max-width: 768px) {
			body {
				padding: 40px 16px;
			}

			.header h1 {
				font-size: 2rem;
			}

			.endpoint-header {
				flex-wrap: wrap;
			}
		}

		::selection {
			background: #d4e5f8;
		}
	</style>
</head>
<body>
	<div class="container">
		<div class="header">
			<h1>API Documentation</h1>
			<div class="meta-info">
				<div class="meta-item">
					<strong>Base URL:</strong>
					<span>{{.BaseURL}}</span>
				</div>
				<div class="meta-item">
					<strong>Version:</strong>
					<span>{{.Version}}</span>
				</div>
				<div class="meta-item">
					<strong>Updated:</strong>
					<span>{{.UpdatedAt}}</span>
				</div>
			</div>
		</div>

		<div class="divider"></div>

		{{range .Endpoints}}
		<div class="endpoint">
			<div class="endpoint-header">
				<span class="method {{.Method}}">{{.Method}}</span>
				<code class="path">{{.Path}}</code>
			</div>
			
			<div class="description">{{.Description}}</div>
			
			{{if .RequestBody}}
			<div class="section">
				<div class="section-title">Request Body</div>
				<pre class="code-block">{{.RequestBody}}</pre>
			</div>
			{{end}}
			
			{{if .ResponseBody}}
			<div class="section">
				<div class="section-title">Response Body</div>
				<pre class="code-block">{{.ResponseBody}}</pre>
			</div>
			{{end}}
			
			<div class="section">
				<div class="section-title">Status Code</div>
				<span class="status-code">{{.StatusCode}}</span>
			</div>
			
			<div class="timestamp">
				<span>Created {{.CreatedAt}}</span>
				<span>Updated {{.UpdatedAt}}</span>
			</div>
		</div>
		{{end}}
	</div>
</body>
</html>
`
