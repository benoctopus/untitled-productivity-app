package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"task-service/ent"
	"task-service/impl"

	pb "proto"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	client, err := ent.Open("postgres", DATABASE_CONNECTION_STRING)
	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
	}

	err = client.Schema.Create(context.Background())
	if err != nil {
		log.Fatalf("failed creating schema to database: %v", err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterTaskServiceServer(s, &impl.TasksServer{Client: client})
	log.Printf("Starting server on port %s", PORT)
	log.Fatal(s.Serve(listener))
}
