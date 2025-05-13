package main

import (
	"context"
	"log"
	"time"

	pb "go-grpc-todo-app-docker/todo"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTodoServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Add a todo
	res, err := client.AddTodo(ctx, &pb.TodoRequest{
		Title:       "Write gRPC app",
		Description: "Build a basic Go gRPC service",
	})
	if err != nil {
		log.Fatalf("AddTodo failed: %v", err)
	}
	log.Printf("Added Todo: %v", res)

	// List todos
	list, err := client.ListTodos(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("ListTodos failed: %v", err)
	}
	for _, t := range list.Todos {
		log.Printf("Todo: %s - %s", t.Title, t.Description)
	}
}
