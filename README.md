# lazydocs

A CLI tool that generates structured API documentation JSON from simple terminal commands.

## Installation
```bash
git clone https://github.com/yourusername/lazydocs.git
cd lazydocs
go build -o lazydocs main.go
```

## Usage
```bash
# Initialize project
./lazydocs -init -v 1.0.0

# Add endpoint
./lazydocs -add api/v1/home -m GET -d "Get homepage" -status 200

# Delete endpoint
./lazydocs -delete api/v1/home -dm GET

# Delete all endpoints
./lazydocs -delete-all
```