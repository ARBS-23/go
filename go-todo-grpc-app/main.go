package main

import (
	"log"
	"net"
	"net/http"

	"github.com/afzalsabbir/go-todo-grpc-app/config"
	pb "github.com/afzalsabbir/go-todo-grpc-app/proto"
	"github.com/afzalsabbir/go-todo-grpc-app/services"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	// Initialize database
	config.ConnectDatabase()

	// Start gRPC server
	go startGRPCServer()

	// Start Gin HTTP server
	startHTTPServer()
}

func startGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTodoServiceServer(grpcServer, &services.TodoServer{})

	log.Printf("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startHTTPServer() {
	r := gin.Default()

	// Define your HTTP endpoints here
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Add more HTTP endpoints as needed

	log.Printf("HTTP server listening on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
} 