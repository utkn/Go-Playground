package main

import (
	pb "../helloworld"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

// Implement the service
func (s *server) SayHello(ctx context.Context, request *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("The server has received a SayHello request.")
	return &pb.GreetResponse{Response: fmt.Sprintf("Hello, dear %v", request.Name)}, nil
}

func (s *server) SayGoodbye(ctx context.Context, request *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("The server has received a SayGoodbye request.")
	return &pb.GreetResponse{Response: fmt.Sprintf("Goodbye, dear %v", request.Name)}, nil
}

// Run the server
func main() {
	// generate a listener on the port 50051
	ls, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer ls.Close()
	// instantiate the grpc server
	grpcServer := grpc.NewServer()
	// register the service on the grpc server
	pb.RegisterGreeterServer(grpcServer, &server{})
	// serve the grpc server on the listener
	err = grpcServer.Serve(ls)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
