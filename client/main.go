package main

import (
	"context"
	"log"
	"strconv"
	"time"

	pb "github.com/smarthobbyjd/go-grpc-postgres/proto/todo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("problem with the server: %v", err)
	}

	defer conn.Close()

	c := pb.NewTodoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	newTodos := []*pb.NewTodo{
		{Name: "Buy pizza", Description: "Make sure to get a spicy one.", Done: false},
		{Name: "Have fun", Description: "Eat and enjoy.", Done: false},
	}

	for _, newTodo := range newTodos {
		res, err := c.CreateTodo(ctx, &pb.NewTodo{Name: newTodo.Name, Description: newTodo.Description, Done: newTodo.Done})
		if err != nil {
			log.Fatalf("could not create todo: %v", err)
		}

		id, err := strconv.ParseInt(res.GetId(), 10, 64)
		if err != nil {
			log.Fatalf("could not convert ID to int64: %v", err)
		}

		log.Printf(`
     ID: %d
     Name: %s
     Description: %s
     Done: %v
    `,
			id,
			res.GetName(),
			res.GetDescription(),
			res.GetDone(),
		)
	}
}
