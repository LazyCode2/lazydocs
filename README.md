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
```bash
# Generates a static site based on JSON
./lazydocs -render 
```
```
```

## 🤝 Contributing
Contributions are welcome! Feel free to open an issue or submit a pull request on Codeberg.
