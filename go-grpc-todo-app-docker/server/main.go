package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	pb "go-grpc-todo-app-docker/todo"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	// Docker compatible PostgreSQL URL
	postgresUrl = "postgres://user:password@postgres:5432/tododb?sslmode=disable"
)

type server struct {
	pb.UnimplementedTodoServiceServer
	db *sql.DB
}

func (s *server) AddTodo(ctx context.Context, req *pb.TodoRequest) (*pb.TodoResponse, error) {
	// Generate unique ID
	id := uuid.New().String()

	// Save to PostgreSQL
	_, err := s.db.Exec("INSERT INTO todos (id, title, description) VALUES ($1, $2, $3)", id, req.Title, req.Description)
	if err != nil {
		return nil, fmt.Errorf("could not insert todo: %v", err)
	}

	// Return response
	return &pb.TodoResponse{
		Id:          id,
		Title:       req.Title,
		Description: req.Description,
	}, nil
}

func (s *server) ListTodos(ctx context.Context, _ *pb.Empty) (*pb.TodoList, error) {
	// Query all todos from PostgreSQL
	rows, err := s.db.Query("SELECT id, title, description FROM todos")
	if err != nil {
		return nil, fmt.Errorf("could not query todos: %v", err)
	}
	defer rows.Close()

	var todos []*pb.TodoResponse
	for rows.Next() {
		var todo pb.TodoResponse
		if err := rows.Scan(&todo.Id, &todo.Title, &todo.Description); err != nil {
			return nil, fmt.Errorf("could not scan todo: %v", err)
		}
		todos = append(todos, &todo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("could not iterate over rows: %v", err)
	}

	return &pb.TodoList{Todos: todos}, nil
}

func main() {
	// Connect to PostgreSQL
	db, err := sql.Open("postgres", postgresUrl)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	// Ensure the "todos" table exists
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS todos (
		id VARCHAR(36) PRIMARY KEY,
		title TEXT,
		description TEXT
	);
	`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	// Set up gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, &server{db: db})
	log.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
