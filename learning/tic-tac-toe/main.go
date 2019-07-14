package main

import "fmt"

const (
	EMPTY Mark = iota //0
	X                 //1
	O                 //2
)

type Mark int
type Board [3][3]Mark

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

func (b Board) String() (out string) {
	for row := 0; row < len(b); row++ {
		for col := 0; col < len(b[row]); col++ {
			cell := b[row][col]
			switch cell {
			case X:
				out += "X"
			case O:
				out += "O"
			case EMPTY:
				out += "_"
			}
		}
		out += "\n"
	}
	return
}

func (b *Board) reset() {
	for row := 0; row < len(b); row++ {
		for col := 0; col < len(b[row]); col++ {
			b[row][col] = EMPTY
		}
	}
}

func validCoords(b Board, row, col int) bool {
	return row >= 0 && row <= 3 && col >= 0 && col <= 3 && b[row][col] == EMPTY
}

func askInput(b Board, turn int) (row, col int) {
	fmt.Printf("Dear %v, please enter input.\n", Mark(turn+1))
	for {
		var line string
		_, _ = fmt.Scanln(&line)
		_, err := fmt.Sscanf(line, "%d,%d", &row, &col)
		if err != nil || !validCoords(b, row, col) {
			fmt.Println("Try again.", err)
		} else {
			break
		}
	}
	return
}

func check(b Board, turn int) bool {
	for i := 0; i < 3; i++ {
		if b[i][0] == Mark(turn+1) && b[i][0] == b[i][1] && b[i][1] == b[i][2] {
			return true
		}
	}
	for i := 0; i < 3; i++ {
		if b[0][i] == Mark(turn+1) && b[0][i] == b[1][i] && b[1][i] == b[2][i] {
			return true
		}
	}
	if b[0][0] == Mark(turn+1) && b[0][0] == b[1][1] && b[1][1] == b[2][2] {
		return true
	}
	if b[0][2] == Mark(turn+1) && b[0][2] == b[1][1] && b[1][1] == b[2][0] {
		return true
	}
	return false
}

func main() {
	turn := 0
	points := [2]int{}
	var b Board
	for {
		fmt.Printf("~~ X:%d ~~ O:%d ~~\n", points[0], points[1])
		fmt.Print(b)
		row, col := askInput(b, turn)
		b[row][col] = Mark(turn + 1)
		if check(b, turn) {
			fmt.Printf("* %d Wins!\n", turn)
			points[turn]++
			b.reset()
		}
		turn = (turn + 1) % 2
	}
}
