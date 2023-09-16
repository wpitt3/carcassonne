package connect4

import (
	"fmt"
	. "will.com/mcts"
)

type ConnectFour struct {
	board  [7][6]int
	player int
}

type ConnectFourAction struct {
	index  int
	player int
}

func NewConnectFour(board [7][6]int, player int) ConnectFour {
	return ConnectFour{
		board,
		player,
	}
}

func (board ConnectFour) Copy() State[Action] {
	var newBoard ConnectFour
	for i := 0; i < 7; i++ {
		for j := 0; j < 6; j++ {
			newBoard.board[i][j] = board.board[i][j]
		}
	}
	newBoard.player = board.player
	return newBoard
}

func (board ConnectFour) ValidActions() []Action {
	validMoves := make([]Action, 0)
	for i := 0; i < 7; i++ {
		if board.board[i][5] == 0 {
			validMoves = append(validMoves, ConnectFourAction{i, board.player})
		}
	}
	return validMoves
}

func (board ConnectFour) IsEndState() bool {
	return boardIsFull(board.board)
}

func (board ConnectFour) Winner() int {
	return find4InBoard(board.board)
}

func (board ConnectFour) CurrentPlayer() int {
	return board.player
}

func (board ConnectFour) PerformMove(action Action) State[Action] {
	connectFourAction := action.(ConnectFourAction)
	for i := 0; i < 6; i++ {
		if board.board[connectFourAction.index][i] == 0 {
			board.board[connectFourAction.index][i] = connectFourAction.player
			board.player = board.player * -1
			return board
		}
	}
	return board
}

func (board ConnectFour) PrintState() {
	valueToPrint := map[int]string{0: " ", 1: "#", -1: "O"}
	var toPrint = ""
	for i := 5; i >= 0; i-- {
		for j := 0; j < 7; j++ {
			toPrint += valueToPrint[board.board[j][i]]
		}
		toPrint += "\n"
	}
	fmt.Println(toPrint)
}

func boardIsFull(board [7][6]int) bool {
	for i := 0; i < 7; i++ {
		if board[i][5] == 0 {
			return false
		}
	}
	return true
}

func find4InBoard(board [7][6]int) int {
	for i := 0; i < 7; i++ {
		for j := 0; j < 3; j++ {
			var sum = 0
			for k := 0; k < 4; k++ {
				sum += board[i][j+k]
			}
			if abs(sum) == 4 {
				return sum / 4
			}
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			var sum = 0
			for k := 0; k < 4; k++ {
				sum += board[i+k][j]
			}
			if abs(sum) == 4 {
				return sum / 4
			}
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			var sumNE = 0
			var sumNW = 0
			for k := 0; k < 4; k++ {
				sumNE += board[i+k][j+k]
				sumNW += board[i+k][3+j-k]
			}
			if abs(sumNE) == 4 {
				return sumNE / 4
			}
			if abs(sumNW) == 4 {
				return sumNW / 4
			}
		}
	}

	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
