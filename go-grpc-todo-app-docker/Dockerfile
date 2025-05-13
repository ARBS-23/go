FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install protobuf compiler for Go
RUN apk add --no-cache protobuf-dev 
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Copy code
COPY . .

# Generate gRPC files
RUN protoc --go_out=. --go-grpc_out=. proto/todo.proto

# Build the server
RUN go mod tidy
RUN go build -o /app/app_server server/main.go
RUN chmod +x /app/app_server

# Use a smaller image for the final container
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/app_server /app/app_server

EXPOSE 50051

# No CMD - will be provided by docker-compose.yml 