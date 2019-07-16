package main

type Mark uint

const (
	EMPTY Mark = iota
	X
	O
)

func (m Mark) String() string {
	switch m {
	case X:
		return "X"
	case O:
		return "O"
	default:
		return "NONE"
	}
}

type GameInstance struct {
	Board
	currentTurn int
	points      [2]int
	names       [2]string
	players     int
	finished    bool
}
