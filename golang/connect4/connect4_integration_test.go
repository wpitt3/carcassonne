package connect4

import (
	"github.com/stretchr/testify/assert"
	"testing"
	. "will.com/mcts"
)

func searchTree() SearchTree {
	return NewSearchTree(RandomRollout{}, float32(1.414), FlatPolicy{})
}

func Test_findBestMove_win(t *testing.T) {

	var board [7][6]int
	board[3][0] = 1
	board[3][1] = 1
	board[3][2] = 1

	assert.Equal(t, ConnectFourAction{3, 1}, searchTree().FindBestMoveByTurns(ConnectFour{board, 1}, 1000))
}

func Test_findBestMove_nearlyLose(t *testing.T) {
	var board [7][6]int
	board[3][0] = -1
	board[3][1] = -1
	board[3][2] = -1

	assert.Equal(t, ConnectFourAction{3, 1}, searchTree().FindBestMoveByTurns(ConnectFour{board, 1}, 1000))
}

func Test_findBestMove_createFork(t *testing.T) {

	//      1
	//      0
	//     00
	//   1 1011
	//   1 0110

	var board [7][6]int
	board[2][0] = 1
	board[2][1] = -1
	board[2][2] = 1
	board[3][0] = -1
	board[3][1] = 1
	board[3][2] = 1
	board[3][3] = 1
	board[3][4] = -1
	board[4][0] = -1
	board[4][1] = -1
	board[5][0] = 1
	board[5][1] = -1
	board[0][0] = -1
	board[0][1] = -1

	assert.Equal(t, ConnectFourAction{4, 1}, searchTree().FindBestMoveByTurns(ConnectFour{board, 1}, 100000))
}
