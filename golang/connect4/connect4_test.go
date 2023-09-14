package connect4

import (
	"github.com/stretchr/testify/assert"
	"testing"
	. "will.com/mcts"
)

func Test_find4_emptyBoard(t *testing.T) {
	var board [7][6]int
	assert.Equal(t, 0, find4InBoard(board))
}

func Test_find4_verticalLine(t *testing.T) {
	var board [7][6]int
	for i := 0; i < 4; i++ {
		board[6][i+2] = 1
	}
	assert.Equal(t, 1, find4InBoard(board))
}

func Test_find4_verticalLineNegative(t *testing.T) {
	var board [7][6]int
	for i := 0; i < 4; i++ {
		board[0][i] = -1
	}
	assert.Equal(t, -1, find4InBoard(board))
}

func Test_find4_horizontalLine(t *testing.T) {
	var board [7][6]int
	for i := 0; i < 4; i++ {
		board[i+3][5] = 1
	}
	assert.Equal(t, 1, find4InBoard(board))
}

func Test_find4_diagNE(t *testing.T) {
	var board [7][6]int
	for i := 0; i < 4; i++ {
		board[3+i][2+i] = 1
	}
	assert.Equal(t, 1, find4InBoard(board))
}

func Test_find4_diagNW(t *testing.T) {
	var board [7][6]int
	for i := 0; i < 4; i++ {
		board[i+3][5-i] = 1
	}
	assert.Equal(t, 1, find4InBoard(board))
}

func Test_copyBoard(t *testing.T) {
	var board ConnectFour
	board.board[0][0] = 1

	var newBoard = board.Copy()
	board.board[0][0] = 0
	assert.Equal(t, 1, newBoard.(ConnectFour).board[0][0])
	assert.Equal(t, 0, board.board[0][0])
}

func Test_fullBoard(t *testing.T) {
	var board [7][6]int
	assert.Equal(t, false, boardIsFull(board))
	for i := 0; i < 7; i++ {
		board[i][5] = 1
	}
	assert.Equal(t, true, boardIsFull(board))
}

func Test_performMove(t *testing.T) {
	var board Board[Action] = ConnectFour{}
	board = board.PerformMove(ConnectFourAction{0, 1})
	board = board.PerformMove(ConnectFourAction{0, 1})
	assert.Equal(t, 1, board.(ConnectFour).board[0][0])
	assert.Equal(t, 1, board.(ConnectFour).board[0][1])
}

func Test_ValidActions(t *testing.T) {
	var board Board[Action] = ConnectFour{}
	actions := board.ValidActions()
	assert.Equal(t, 7, len(actions))
}