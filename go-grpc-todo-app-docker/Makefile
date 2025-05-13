.PHONY: proto run-server run-client docker-build docker-run

proto:
	protoc --go_out=. --go-grpc_out=. proto/todo.proto

run-server:
	go run server/main.go

run-client:
	go run client/main.go

docker-build:
	docker build -t go-grpc-todo .

docker-run:
	docker run -p 50051:50051 go-grpc-todo
