package main

import (
	pb "../game"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"strings"
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println("Welcome to online Tic tac toe.")

	err = errors.New("")
	var playerName string
	for err != nil {
		fmt.Println(err)
		fmt.Print("Please enter a name: ")
		_, err = fmt.Scanln(&playerName)
	}

	joinRsp, err := stb.JoinRandomGame(ctx, &pb.JoinRequest{
		Name: playerName,
	})

	if err != nil {
		log.Fatal(err)
	}

	if !joinRsp.Success {
		for {
			fmt.Println("Waiting for another player to join...")
			time.Sleep(time.Second * 5)
			gameState, err := stb.GetGameState(ctx, joinRsp.PlayerId)
			if err != nil {
				log.Fatal(err)
			}
			if gameState.Turn {
				fmt.Println("Someone has joined!")
				break
			}
		}
	}

	for {
		gameState, err := stb.GetGameState(ctx, joinRsp.PlayerId)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(gameState.Cells)
		if gameState.Finished && gameState.Turn {
			fmt.Println("You won!")
			break
		} else if gameState.Finished {
			fmt.Println("You lost!")
			break
		} else if gameState.Turn {
			fmt.Print("It's your turn: ")
			var in string
			fmt.Scanln(&in)
			coords := strings.Split(in, ",")
			coordsInt := [2]int{}
			coordsInt[0], _ = strconv.Atoi(coords[0])
			coordsInt[1], _ = strconv.Atoi(coords[1])
			stb.PutMark(ctx, &pb.Action{
				PlayerId: joinRsp.PlayerId,
				Cell:     int32(coordsInt[0]*3 + coordsInt[1]),
			})
		} else {
			fmt.Println("Waiting for the opposing player...")
			time.Sleep(time.Second * 5)
		}
	}
}
