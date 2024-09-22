# Golang REST API with Repository Pattern and WebSocket

This project is a basic implementation of a **REST API** in Go using the **Repository Pattern** to manage data access and a **WebSocket** for real-time communication.

## Table of Contents
- [Features](#features)
- [Requirements](#requirements)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [WebSocket](#websocket)
- [Repository Pattern](#repository-pattern)
## Features
- REST API for managing resources.
- Repository Pattern to abstract database logic.
- WebSocket support for real-time updates.
- Modular and scalable project structure.

## Requirements
- Go 1.23.0+
- Gorilla WebSocket (`github.com/gorilla/websocket`)
- Mux router (`github.com/gorilla/mux`)
-Others
## Project Structure

```
.
├── databases
│   └── postgress.go        # Query and conection with DB
├── handlers
│   └── user.go             # Controller for handling HTTP requests
├── models
│   └── user.go             # User model definition
├── repository
│   └── repository.go       # Repository pattern implementation
├── websocket
│   └── websocket.go        # WebSocket implementation
├── go.mod                  # Go module file
├── go.sum                  # Go dependencies
└── README.md               # Project documentation
```

## Installation

### Clone the repository:

```bash
git clone https://github.com/mariasilva795/go-api-rest.git
cd your-repository
```

### Install dependencies:

```bash
go mod tidy
```

## Usage

1. **Run the application**:

   To run the Go application, use:

   ```bash
   go run main.go
   ```

2. **Testing the REST API**:

   You can test the REST API endpoints using tools like [Postman](https://www.postman.com/) or [curl](https://curl.se/).

3. **Testing the WebSocket**:

   You can test the WebSocket functionality using a WebSocket client or browser developer tools.

## WebSocket

The WebSocket API makes it possible to open a two-way interactive communication session between the user's browser and a server. With this API, you can send messages to a server and receive responses without having to poll the server for a reply.


### WebSocket Endpoint

| Endpoint  | Description       |
|-----------|-------------------|
| `/ws`     | WebSocket handler |

#### Example JavaScript Client:

```javascript
    var ws = new WebSocket("ws://localhost:5050/ws");

    ws.onopen = function() {
      console.log("Connected to server");
    };

    ws.onmessage = function(event) {
      console.log("Received message: " + event.data);
    };

    ws.onerror = function(event) {
      console.log("Error: " + event.data);
    };

    fetch("http://localhost:5050/posts", {
      method: "GET",
      headers: {
        "Content-Type": "application/json"
      }
    }).then(function(response) {
      return response.json();
    }).then(function(json) {
      console.log('DATA RESPONSE')
      console.log(json);
    });
```

## Repository Pattern

The Repository Pattern is used to separate the logic of data retrieval and storage from the business logic. This makes the code more maintainable and testable by abstracting the data layer.

### Example Implementation:

**UserRepository Interface:**

```go
type UserRepository interface {
    GetAll() ([]User, error)
    GetByID(id string) (User, error)
    Create(user User) error
    Update(id string, user User) error
    Delete(id string) error
}
```

**UserService:**

```go
type UserService struct {
    repo UserRepository
}

func (s *UserService) GetAllUsers() ([]User, error) {
    return s.repo.GetAll()
}
```