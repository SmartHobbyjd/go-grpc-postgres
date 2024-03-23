package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"

	pb "github.com/smarthobbyjd/go-grpc-postgres/proto/todo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"google.golang.org/grpc"
)

type TodoServer struct {
	pb.UnimplementedTodoServiceServer
}

func init() {
	DatabaseConnection()
}

var DB *gorm.DB
var err error

func DatabaseConnection() {
	DB, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal("db connection error: ", err)
	}
	log.Println("db connection successful")
}

func (s *TodoServer) CreateTodo(ctx context.Context, req *pb.NewTodo) (*pb.Todo, error) {
	log.Printf("Received: %v", req.GetName())
	todo := &pb.Todo{
		Name:        req.GetName(),
		Description: req.GetDescription(),
	}
	res := DB.Create(&todo)
	if res.RowsAffected == 0 {
		return nil, errors.New("error saving todo")
	}
	return &pb.Todo{
		Id:          todo.Id,
		Name:        todo.Name,
		Description: todo.Description,
		Done:        false,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("tcp connection failed: %v", err)
	}
	log.Printf("listening at %v", lis.Addr())

	s := grpc.NewServer()

	pb.RegisterTodoServiceServer(s, &TodoServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("grpc server failed: %v", err)
	}
}
