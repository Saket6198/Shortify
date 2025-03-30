# GoShortify - A Golang Powered Url Shortner

![Go](https://img.shields.io/badge/Go-1.20%2B-00ADD8?logo=go&logoColor=white) 
![License](https://img.shields.io/badge/License-MIT-blue) 

A simple URL shortening service built in Go. This project generates concise URLs for long web addresses and redirects users from the shortened URL to the original URL using the **latest and industry grade SHA-256 Encryption for checksums which allows for 0% chance of collision unlike previous methods**. It demonstrates a basic REST API design, modular file structure, error handling, and logging.

## Table of Contents
- [Project Structure](#project-structure)
- [Features](#Features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Data Model](#data-model)
- [System Design Overview](#system-design-overview)
- [Error Handling & Logging](#error-handling--logging)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgements](#acknowledgements)

## File Structure
```
url-shortener/
├── handlers/
│   ├── shorten.go       # Handles POST /shorten request
│   └── redirect.go      # Handles GET /redirect/{id} request
├── models/
│   └── url.go           # Contains the Url struct and in-memory storage
├── routes/
│   └── routes.go        # Registers HTTP routes and associates endpoints with handlers
├── utils/
│   └── hash.go          # Contains utility functions for hashing URLs
├── main.go              # Application entry point
└── go.mod               # Go module file
```


## Features

- **URL Shortening**: Generate a unique short URL for any provided long URL.
- **Redirection**: Redirect users from the short URL to the original URL.
- **REST API**: Exposes endpoints for URL creation and redirection.
- **Modular Design**: Well-organized file structure with separate packages for models, handlers, utilities, and routes.
- **Error Handling**: Comprehensive error responses for invalid input and missing URLs.
- **Logging**: Logs significant events such as redirections and errors.

## Prerequisites

- [Go](https://golang.org/dl/) version 1.16 or higher

## Installation

1. **Clone the Repository:**

    ```sh
    git clone https://github.com/yourusername/url-shortener.git
    cd url-shortener
    ```

2. **Initialize the Go Module:**

    ```sh
    go mod init url-shortener
    go mod tidy
    ```

## Usage

1. **Run the Application:**

    ```sh
    go run main.go
    ```

2. **Access the Service:**
   - The server starts on port `5000`.
   - Use your preferred API client (like Postman or cURL) to interact with the endpoints.

## API Endpoints

### POST /shorten

- **Description**: Shortens a provided URL.
- **Request Body**:

    ```json
    {
      "url": "http://example.com"
    }
    ```
![image](https://github.com/user-attachments/assets/e7142f03-f7db-42da-9a6f-b22482d68590)

- **Response**:

    ```json
    {
      "shortened_url": "shortUrlID"
    }
    ```

### GET /redirect/{id}

- **Description**: Redirects to the original URL associated with the given short URL ID.
- **URL Parameter**:
  - `{id}`: The identifier for the shortened URL.
- **Response**: The service redirects to the original URL.
- ![image](https://github.com/user-attachments/assets/5f20e2c4-5113-46c8-bc35-7fa16418151b)


## Data Model

The primary data model used for URL storage is as follows:

```go
type Url struct {
    Id           string    `json:"id"`
    OriginalUrl  string    `json:"original_url"`
    ShortenedUrl string    `json:"shortened_url"`
    CreationDate time.Time `json:"creation_date"`
}
```
- *Id*: Unique identifier generated via SHA-256 hash (first 16 characters).

- *OriginalUrl*: The long URL provided by the user.

- *ShortenedUrl*: The generated short URL.

- *CreationDate*: Timestamp when the URL was shortened.

## System Design Overview

1. Modular Architecture:
  The project is structured into distinct packages to separate concerns:
  
    - Models: Contains data structures and in-memory storage.
    
    - Handlers: Implements business logic for handling HTTP requests.
    
    - Utils: Provides utility functions (e.g., SHA-256 hashing).
    
    - Routes: Centralizes route registration for clean API definitions.

2. REST API Design:
  The API exposes two main endpoints:
  
    - POST /shorten: Accepts a JSON payload to create a shortened URL.
    
    - GET /redirect/{id}: Redirects to the original URL using the provided short URL ID.

3. Scalability Considerations:
  While the current implementation uses in-memory storage (making it volatile), the modular design allows for easy integration with persistent storage (e.g., SQL or NoSQL databases) when needed.

4. Hashing Strategy:
  The service uses SHA-256 to generate a hash for the given URL and truncates it to create a unique identifier. This minimizes the risk of collisions while keeping the URL concise.

## Error Handling & Logging
- Error Handling:

  - Invalid JSON input during URL shortening returns a 400 Bad Request error.

  - Requests for non-existent URLs return a 404 Not Found error.

- Logging:

 - The application logs key events, such as successful redirection and errors, using Go's standard log package.

## Contributing
- Fork the repo

- Create your feature branch

- Commit changes

- Push and open a PR

---
