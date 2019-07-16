package main

import (
	pb "../helloworld"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"sync"
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create two concurrent users sending two requests synchronously to test how the server/client behaves.
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func(z int) {
			// first request
			r, err := stb.SayHello(ctx, &pb.GreetRequest{Name: fmt.Sprint("Utkan", z)})
			if err != nil {
				log.Fatalf("Could not send the request: %v", err)
			}
			log.Printf("Acquired: %v", r.Response)
			// second request
			r, err = stb.SayGoodbye(ctx, &pb.GreetRequest{Name: fmt.Sprint("Utkan", z)})
			if err != nil {
				log.Fatalf("Could not send the request: %v", err)
			}
			log.Printf("Acquired: %v", r.Response)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
