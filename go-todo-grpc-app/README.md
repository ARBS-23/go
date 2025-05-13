# Go Todo gRPC Application

This is a Todo application built with Go, using Gin for HTTP API, GORM for database operations, and gRPC for service communication.

## Features

- gRPC server for Todo operations
- HTTP server using Gin framework
- SQLite database with GORM
- Full CRUD operations for Todo items

## Prerequisites

- Go 1.21 or later
- Protocol Buffers compiler (protoc)
- Go gRPC tools

## Installation

1. Clone the repository:
```bash
git clone https://github.com/afzalsabbir/go-todo-grpc-app.git
cd go-todo-grpc-app
```

2. Install dependencies:
```bash
go mod tidy
```

3. Generate Protocol Buffers code:
```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/todo.proto
```

## Running the Application

1. Start the application:
```bash
go run main.go
```

This will start:
- gRPC server on port 50051
- HTTP server on port 8080
- SQLite database will be created automatically

## API Endpoints

### gRPC Endpoints

The following gRPC services are available:

- CreateTodo
- GetTodo
- ListTodos
- UpdateTodo
- DeleteTodo

### HTTP Endpoints

- GET /ping - Health check endpoint

## Testing

You can use tools like [grpcurl](https://github.com/fullstorydev/grpcurl) to test the gRPC endpoints:

```bash
# List services
grpcurl -plaintext localhost:50051 list

# Create a todo
grpcurl -plaintext -d '{"title": "Test Todo", "description": "Test Description", "completed": false}' \
    localhost:50051 proto.TodoService/CreateTodo

# List todos
grpcurl -plaintext localhost:50051 proto.TodoService/ListTodos
```

## License

MIT 