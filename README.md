# lazydocs

A CLI tool that generates structured API documentation JSON from simple terminal commands.

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
[![License](https://img.shields.io/badge/license-MIT-green)](https://codeberg.org/LazyCode2/lazydocs/src/branch/main/LICENSE)
![Status](https://img.shields.io/badge/status-active-success)
![Platform](https://img.shields.io/badge/platform-linux%20%7C%20macos%20%7C%20windows-lightgrey)

## 🛠 Installation
```bash
git clone https://codeberg.org/LazyCode2/lazydocs.git
cd lazydocs
go build -o lazydocs main.go
```

## 📖 Usage

#### Initialize a Project :
Start by creating a new documentation project:

```Bash
./lazydocs -init -v 1.2.0
```

#### Adding Endpoints :
Add a new API endpoint with descriptions and status codes:

```Bash
./lazydocs -add api/v1/users -m POST -d "Create a new user" -status 201

# Full endpoint with body
./lazydocs -add api/v1/users -m POST \
  -d "Create user" \
  -body '{"name": "string", "email": "string"}' \
  -response '{"id": "string", "name": "string"}' \
  -status 201
```
#### Listing Endpoints :
View your documented endpoints:

```Bash
# List all endpoints
./lazydocs -list

# List only GET endpoints
./lazydocs -list GET
```

#### Deleting Endpoints :
Remove specific endpoints or clear the entire project:

```Bash
# Delete a specific endpoint
./lazydocs -delete api/v1/users -dm POST

# Delete all
./lazydocs -delete-all
```

## Site rendering
Render your JSON documentation as a static HTML site:
```bash
# Generates a static site based on JSON
./lazydocs -render 
```


This generates a clean, styled documentation site in the `docs/` folder based on your documented endpoints.

## Template customization

After running `-init`, lazydocs creates a default template at `template/index.tmpl`. You can customize this template to match your branding or documentation style.

### Available template variables

The template has access to the following data:

- `{{.BaseURL}}` - Your API's base URL
- `{{.Version}}` - API version number
- `{{.UpdatedAt}}` - Last update timestamp
- `{{.Endpoints}}` - Array of all documented endpoints

Each endpoint in `{{.Endpoints}}` contains:

- `{{.Method}}` - HTTP method (GET, POST, etc.)
- `{{.Path}}` - Endpoint path
- `{{.Description}}` - Endpoint description
- `{{.RequestBody}}` - Request body schema (if provided)
- `{{.ResponseBody}}` - Response body schema (if provided)
- `{{.StatusCode}}` - HTTP status code
- `{{.CreatedAt}}` - Creation timestamp
- `{{.UpdatedAt}}` - Last update timestamp

### Example customization

You can modify `template/index.tmpl` to change colors, fonts, layout, or add additional sections. The template uses Go's `html/template` syntax, so you can add conditionals, loops, and custom formatting.

For instance, to add a dark theme, modify the CSS section in the template. To change the layout structure, rearrange the HTML elements. Any changes you make will be reflected when you run `./lazydocs -render`.

## 🤝 Contributing
Contributions are welcome! Feel free to open an issue or submit a pull request on Codeberg.

## Command reference
```
Project:
  -init                    Initialize project
  -v <version>             API version (default: 1.0.0)
  -base <base_url>         Set base URL

View:
  -list                    List all API endpoints
  -lm <method>             List endpoints by HTTP method

Site:
  -render                  Generate documentation site from JSON

Add Endpoint:
  -add <endpoint>          Add new API endpoint
  -m <method>              HTTP method (default: GET)
  -d <description>         Endpoint description
  -body <schema>           Request body (optional)
  -response <schema>       Response body (optional)
  -status <code>           HTTP status code (default: 200)

Delete:
  -delete <endpoint>       Delete specific endpoint
  -dm <method>             Method of endpoint to delete
  -delete-all              Delete all endpoints

Other:
  -h                       Show help
```
