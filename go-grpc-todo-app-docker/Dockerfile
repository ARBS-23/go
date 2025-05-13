FROM golang:1.23-alpine

WORKDIR /app

# Copy files
COPY . .

# Install protobuf compiler for Go
RUN apk add --no-cache protobuf-dev 
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate gRPC files
RUN mkdir -p todo
RUN protoc --go_out=. --go-grpc_out=. proto/todo.proto

# Build the server
RUN go mod tidy
RUN cd server && go build -o main

EXPOSE 50051

CMD ["./server/main"]
