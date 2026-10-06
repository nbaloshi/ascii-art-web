# ASCII Art Web

A web application written in Go that converts user-provided text into ASCII art and displays the result through a simple web interface.

The project started as a command-line ASCII art generator and was extended into a web application with HTML templates, form handling, custom error pages, CSS styling, and Docker support.

## Features

* Convert text into ASCII art
* Web-based interface
* HTML templates
* Form handling with POST requests
* Input validation
* Multiple ASCII-art banner styles
* Custom HTTP error pages
* CSS styling
* Docker support

## Tech Stack

* Go
* HTML
* CSS
* Go `html/template`
* Go `net/http`
* Docker

## How It Works

The Go HTTP server listens on port `8080`.

Users submit text through the web interface. The server:

1. Receives the submitted form data.
2. Validates the input.
3. Loads the selected ASCII-art banner.
4. Converts each character into its corresponding ASCII-art representation.
5. Renders the result back through the HTML template.

## Project Structure

```text
.
├── main.go
├── go.mod
├── standard.txt
├── thinkertoy2.txt
├── web.html
├── web.css
├── error400.html
├── error404.html
├── error405.html
├── error500.html
└── dockerfile
```

## Getting Started

### Prerequisites

* Go
* Docker (optional)

### Run Locally

Clone the repository and start the server:

```bash
go run .
```

The application will be available at:

```text
http://localhost:8080
```

### Build

```bash
go build
```

### Docker

Build the Docker image:

```bash
docker build -f dockerfile -t ascii-art-web .
```

Run the container:

```bash
docker run -p 8080:8080 ascii-art-web
```

Then open:

```text
http://localhost:8080
```

## Project Status

This project is a completed academic project from my full-stack software development studies.

It demonstrates foundational experience with Go web servers, HTTP request handling, HTML templates, form processing, input validation, file handling, and containerization.
