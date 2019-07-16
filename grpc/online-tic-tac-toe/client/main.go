package main

import (
	pb "../game"
	"Go-Playground/grpc/online-tic-tac-toe/game"
	"context"
	"errors"
	"fmt"
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
	stb := pb.NewGameClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("Welcome to online Tic tac toe.")

	err = errors.New("")
	var playerName string
	for err != nil {
		fmt.Println(err)
		fmt.Print("Please enter a name: ")
		_, err = fmt.Scanln(&playerName)
	}

	_, err = stb.JoinRandomGame(ctx, &game.JoinRequest{
		Name: playerName,
	})

	if err != nil {
		log.Fatal(err)
	}

}
