package services

import (
	"context"

	"github.com/afzalsabbir/go-todo-grpc-app/config"
	"github.com/afzalsabbir/go-todo-grpc-app/models"
	pb "github.com/afzalsabbir/go-todo-grpc-app/proto"
)

type TodoServer struct {
	pb.UnimplementedTodoServiceServer
}

func (s *TodoServer) CreateTodo(ctx context.Context, req *pb.Todo) (*pb.Todo, error) {
	todo := &models.Todo{
		Title:       req.Title,
		Description: req.Description,
		Completed:   req.Completed,
	}

	result := config.DB.Create(todo)
	if result.Error != nil {
		return nil, result.Error
	}

	return &pb.Todo{
		Id:          uint64(todo.ID),
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
	}, nil
}

func (s *TodoServer) GetTodo(ctx context.Context, req *pb.TodoId) (*pb.Todo, error) {
	var todo models.Todo
	result := config.DB.First(&todo, req.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &pb.Todo{
		Id:          uint64(todo.ID),
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
	}, nil
}

func (s *TodoServer) ListTodos(ctx context.Context, req *pb.Empty) (*pb.TodoList, error) {
	var todos []models.Todo
	result := config.DB.Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}

	var pbTodos []*pb.Todo
	for _, todo := range todos {
		pbTodos = append(pbTodos, &pb.Todo{
			Id:          uint64(todo.ID),
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   todo.Completed,
		})
	}

	return &pb.TodoList{Todos: pbTodos}, nil
}

func (s *TodoServer) UpdateTodo(ctx context.Context, req *pb.Todo) (*pb.Todo, error) {
	var todo models.Todo
	result := config.DB.First(&todo, req.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	todo.Title = req.Title
	todo.Description = req.Description
	todo.Completed = req.Completed

	result = config.DB.Save(&todo)
	if result.Error != nil {
		return nil, result.Error
	}

	return &pb.Todo{
		Id:          uint64(todo.ID),
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
	}, nil
}

func (s *TodoServer) DeleteTodo(ctx context.Context, req *pb.TodoId) (*pb.Empty, error) {
	result := config.DB.Delete(&models.Todo{}, req.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &pb.Empty{}, nil
} 