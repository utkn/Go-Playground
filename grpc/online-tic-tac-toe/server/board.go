package main

type Board [3][3]Mark

func checkBoard(b Board, turn int) bool {
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
