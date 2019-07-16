package main

import (
	pb "Go-Playground/grpc/online-tic-tac-toe/game"
	"context"
	"errors"
	"log"
	"sync"
)

type Server struct {
	lastId        int
	GameInstances map[int](*GameInstance)
	GameMutex     *sync.Mutex
}

func (s *Server) newGame() (*GameInstance, int) {
	s.GameInstances[s.lastId] = &GameInstance{}
	s.lastId++
	return s.GameInstances[s.lastId-1], s.lastId - 1
}

func (s *Server) findAvailableGame() (*GameInstance, int) {
	for k, v := range s.GameInstances {
		if v.players < 2 {
			return v, k
		}
	}
	return nil, 0
}

// implement the game service
func (s *Server) JoinRandomGame(ctx context.Context, req *pb.JoinRequest) (*pb.JoinResponse, error) {
	log.Printf("Join request received from %v\n", req.Name)
	s.GameMutex.Lock()
	defer s.GameMutex.Unlock()

	newGame := false
	player := 1
	availableGame, gameId := s.findAvailableGame()
	if availableGame == nil {
		newGame = true
		player = 0
		availableGame, gameId = s.newGame()
	}

	availableGame.names[player] = req.Name
	availableGame.players = availableGame.players + 1

	return &pb.JoinResponse{
		Success: !newGame,
		PlayerId: &pb.PlayerId{
			GameId: int32(gameId),
			Player: int32(player),
		}}, nil
}

func (s *Server) GetGameState(ctx context.Context, req *pb.PlayerId) (*pb.GameState, error) {
	log.Printf("Game state request received from %v\n", req.Player)
	s.GameMutex.Lock()
	defer s.GameMutex.Unlock()
	gameInstance, ok := s.GameInstances[int(req.GameId)]
	if !ok {
		return nil, errors.New("could not find game")
	}

	cells := [9]uint32{}

	for i := range gameInstance.Board {
		for j := range gameInstance.Board[i] {
			cells[i*3+j] = uint32(gameInstance.Board[i][j])
		}
	}

	return &pb.GameState{
		Finished: gameInstance.finished,
		Turn:     gameInstance.players == 2 && gameInstance.currentTurn == int(req.Player),
		Cells:    cells[:],
	}, nil

}

func (s *Server) PutMark(ctx context.Context, req *pb.Action) (*pb.GameState, error) {
	log.Printf("Put mark request received from %v\n", req.PlayerId)
	s.GameMutex.Lock()
	defer s.GameMutex.Unlock()
	gameInstance, ok := s.GameInstances[int(req.PlayerId.GameId)]
	if !ok {
		return nil, errors.New("could not find game")
	} else if gameInstance.players != 2 {
		return nil, errors.New("the game has not started yet")
	} else if gameInstance.currentTurn != int(req.PlayerId.Player) {
		return nil, errors.New("it's not your turn")
	}

	cells := [9]uint32{}

	for i := range gameInstance.Board {
		for j := range gameInstance.Board[i] {
			cells[i*3+j] = uint32(gameInstance.Board[i][j])
		}
	}

	row := req.Cell / 3
	col := req.Cell % 3
	gameInstance.Board[row][col] = Mark(req.PlayerId.Player + 1)
	if checkBoard(gameInstance.Board, gameInstance.currentTurn) {
		gameInstance.finished = true
	}
	gameInstance.currentTurn = (gameInstance.currentTurn + 1) % 2

	return &pb.GameState{
		Finished: gameInstance.finished,
		Turn:     gameInstance.currentTurn == int(req.PlayerId.Player),
		Cells:    cells[:]}, nil
}
