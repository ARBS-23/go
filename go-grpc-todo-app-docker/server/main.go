package main

import (
	"context"
	"log"
	"net"

	pb "go-grpc-todo-app-docker/todo"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTodoServiceServer
	todos []*pb.TodoResponse
}

func (s *server) AddTodo(ctx context.Context, req *pb.TodoRequest) (*pb.TodoResponse, error) {
	todo := &pb.TodoResponse{
		Id:          uuid.New().String(),
		Title:       req.Title,
		Description: req.Description,
	}
	s.todos = append(s.todos, todo)
	return todo, nil
}

func (s *server) ListTodos(ctx context.Context, _ *pb.Empty) (*pb.TodoList, error) {
	return &pb.TodoList{Todos: s.todos}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, &server{})
	log.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
