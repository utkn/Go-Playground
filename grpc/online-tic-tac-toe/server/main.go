package main

import (
	pb "Go-Playground/grpc/online-tic-tac-toe/game"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

func main() {
	// generate a listener on the port 50051
	ls, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer ls.Close()
	// instantiate the gRPC server
	grpcServer := grpc.NewServer()
	// register the service on the gRPC server
	gameServer := Server{
		GameInstances: make(map[int]*GameInstance),
		GameMutex:     &sync.Mutex{},
	}

	go func() {
		for {
			var command string
			fmt.Print("Enter command: ")
			fmt.Scanln(&command)
			if command == "1" {
				gameServer.GameMutex.Lock()
				fmt.Println(gameServer.GameInstances[0])
				gameServer.GameMutex.Unlock()
			}
		}
	}()

	pb.RegisterGameServer(grpcServer, &gameServer)
	// serve the gRPC server on the listener
	err = grpcServer.Serve(ls)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
