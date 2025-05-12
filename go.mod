module go-grpc-todo-app-docker

go 1.21

require (
	github.com/google/uuid v1.6.0
	google.golang.org/grpc v1.62.0
	google.golang.org/protobuf v1.34.1
)

replace todo => ./todo
