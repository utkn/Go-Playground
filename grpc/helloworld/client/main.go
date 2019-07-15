package main

import (
	pb "../helloworld"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	// stub
	stb := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := stb.SayHello(ctx, &pb.GreetRequest{Name: "Utkan"})
	if err != nil {
		log.Fatalf("Could not send the request: %v", err)
	}

	log.Printf("Acquired: %v", r.Response)

	r, err = stb.SayGoodbye(ctx, &pb.GreetRequest{Name: "Utkan"})
	if err != nil {
		log.Fatalf("Could not send the request: %v", err)
	}

	log.Printf("Acquired: %v", r.Response)
}
