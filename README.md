# GoShortify - A Golang and Redis Powered Url Shortner

![Go](https://img.shields.io/badge/Go-1.20%2B-00ADD8?logo=go&logoColor=white) 
![License](https://img.shields.io/badge/License-MIT-blue) 

![image](https://github.com/user-attachments/assets/71ff80ff-6b79-4aa1-922b-fab7e6104839)


A URL shortening service built in Go that generates concise URLs for long web addresses and utilizes **Redis for low-latency caching and ultra-fast performance**. It redirects users from the shortened URL to the original URL using the **latest industry-grade SHA-256 encryption for checksums, ensuring a 0% chance of collision compared to previous methods**. The service demonstrates a basic REST API design, a modular file structure, robust error handling, and comprehensive logging.

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


## File Structure
```
url-shortener/
├── handlers/
│   ├── shorten.go       # Handles POST /shorten request
│   ├── redirect.go      # Handles GET /redirect/{id} request
│   ├── clear.go         # Handles DELETE /clear request (flush Redis)
├── models/
│   └── url.go           # Contains the Url struct and in-memory storage
├── routes/
│   └── routes.go        # Registers HTTP routes and associates endpoints with handlers
├── utils/
│   ├── hash.go          # Contains utility functions for hashing URLs
│   ├── redis.go         # Initializes and manages Redis connection
├── main.go              # Application entry point
└── go.mod               # Go module file
```


## Features

- **URL Shortening**: Generate a unique short URL for any provided long URL.
- **SHA-256 Hashing**: Ensures 0% chance of collision for short URLs.
- **Redirection**: Redirect users from the short URL to the original URL.
- **REST API**: Exposes endpoints for URL creation, redirection, and clearing data.
- **Modular Design**: Well-organized file structure with separate packages for models, handlers, utilities, and routes.
- **Error Handling**: Comprehensive error responses for invalid input and missing URLs.
- **Logging**: Logs significant events such as redirections, errors, and Redis operations.

## Prerequisites

- [Go](https://golang.org/dl/) version 1.16 or higher
- [Redis](redis.io) installed and running

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
    
3. **Start Redis:**
   
   ![image](https://github.com/user-attachments/assets/b3ef0619-d3c0-406e-921b-1ad0a3d9f464)

   If you are using WSL on windows or linux , install and start redis with:
    ```bash
    sudo apt-get install redis-server
    redis-server
    ```
    Ensure the port 6379(default by redis) is open, else run
   ```bash
   netstat -aonb | findstr 6379
   ```
   and kill the pid and restart redis.    
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

### DELETE /clear

- **Description:** Clears all stored URL mappings from Redis.

- **Response**:
    - Success: 200 OK - "Cleared all data from Redis"

    - Failure: 500 Internal Server Error

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


3. Redis for Storage:
URLs are stored in Redis for fast lookup and persistence.
    - Redis SET and GET commands retrieve URLs quickly.

    - The /clear route flushes all stored URLs when needed.


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
