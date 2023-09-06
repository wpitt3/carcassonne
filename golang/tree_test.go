package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func shite(t *testing.T) {
	fmt.Printf("Go says %v\n", 0)
}

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

func Test_rootnode(t *testing.T) {
	var board Board[Action] = ConnectFour{[7][6]int{}, 1}
	board = board.PerformMove(ConnectFourAction{0, 1}).(ConnectFour)
	rootNode := rootNode(board)

	assert.Equal(t, nil, rootNode.action)
	assert.Equal(t, 1, rootNode.board.(ConnectFour).board[0][0])
	assert.Equal(t, -1, rootNode.board.(ConnectFour).player)
	assert.Equal(t, 7, len(rootNode.unexpandedActions))
}

func Test_newNode(t *testing.T) {
	var board Board[Action] = ConnectFour{[7][6]int{}, 1}
	rootNode := rootNode(board)
	newNode := newNode(rootNode, ConnectFourAction{0, 1})

	assert.Equal(t, ConnectFourAction{0, 1}, newNode.action)
	assert.Equal(t, 1, newNode.board.(ConnectFour).board[0][0])
	assert.Equal(t, -1, newNode.board.(ConnectFour).player)
	assert.Equal(t, rootNode, newNode.parent)
	assert.Equal(t, 0, newNode.winner)
	assert.Equal(t, false, newNode.isTerminalState)
	assert.Equal(t, 7, len(rootNode.unexpandedActions))
	assert.Equal(t, rootNode.children[0], newNode)
}

func Test_newNodeHasWinner(t *testing.T) {
	var board Board[Action] = ConnectFour{[7][6]int{}, 1}
	board = board.PerformMove(ConnectFourAction{0, 1})
	board = board.PerformMove(ConnectFourAction{0, 1})
	board = board.PerformMove(ConnectFourAction{0, 1})
	rootNode := rootNode(board)
	newNode := newNode(rootNode, ConnectFourAction{0, 1})

	assert.Equal(t, ConnectFourAction{0, 1}, newNode.action)
	assert.Equal(t, 1, newNode.winner)
	assert.Equal(t, true, newNode.isTerminalState)
}

func Test_newNodeIsFullWithNoWinner(t *testing.T) {
	var board [7][6]int
	for i := 0; i < 5; i++ {
		board[0][i] = 1
	}
	for i := 1; i < 7; i++ {
		board[i][5] = 1
	}
	board[0][3] = -1
	board[3][5] = -1
	rootNode := rootNode(ConnectFour{board, 1})
	newNode := newNode(rootNode, ConnectFourAction{0, 1})

	assert.Equal(t, ConnectFourAction{0, 1}, newNode.action)
	assert.Equal(t, 0, newNode.winner)
	assert.Equal(t, true, newNode.isTerminalState)
}

func Test_findBestChild(t *testing.T) {
	var board Board[Action] = ConnectFour{[7][6]int{}, 1}
	rootNode := rootNode(board)
	nodeA := newNode(rootNode, ConnectFourAction{1, 1})
	nodeB := newNode(rootNode, ConnectFourAction{2, 1})

	rootNode.numer = 1.0
	rootNode.denom = 4.0
	nodeA.numer = 1.0
	nodeA.denom = 2.0
	nodeB.numer = 2.0
	nodeB.denom = 2.0

	// B is better move and should be explored more
	assert.Equal(t, nodeB, findBestChild(rootNode, float32(1.414)))

	// B is better move still, but a should be explored more as b is overly explored
	rootNode.denom = 5.0
	nodeB.denom = 3.0
	assert.Equal(t, nodeA, findBestChild(rootNode, float32(1.414)))
}

func Test_selectLeafNode_expandRootNode(t *testing.T) {
	var board Board[Action] = ConnectFour{[7][6]int{}, 1}
	rootNode := rootNode(board)
	action1 := rootNode.unexpandedActions[0]

	node := selectLeafNode(rootNode)
	assert.Equal(t, node.action, action1)
}

//
//func Test_selectLeafNode_expandChild(t *testing.T) {
//	var board ConnectFour
//	rootNode := rootNode(board)
//	rootNode.unexpandedActions = []ConnectFourAction
//	nodeA := newNode(rootNode, 1)
//	action1 := nodeA.unexpandedActions[0]
//
//	node := selectLeafNode(rootNode)
//	assert.Equal(t, node.action, action1)
//}
//
//func Test_selectLeafNode_dontExpandTerminal(t *testing.T) {
//	var board ConnectFour
//	rootNode := rootNode(board)
//	rootNode.unexpandedActions = []ConnectFourAction
//	nodeA := newNode(rootNode, 1)
//	nodeA.isTerminalState = true
//
//	node := selectLeafNode(rootNode)
//	assert.Equal(t, node, nodeA)
//}

func Test_findBestMove_win(t *testing.T) {
	var board [7][6]int
	board[3][0] = 1
	board[3][1] = 1
	board[3][2] = 1

	assert.Equal(t, ConnectFourAction{3, 1}, findBestMove(ConnectFour{board, 1}, 1000))
}

func Test_findBestMove_nearlyLose(t *testing.T) {
	var board [7][6]int
	board[3][0] = -1
	board[3][1] = -1
	board[3][2] = -1

	assert.Equal(t, ConnectFourAction{3, 1}, findBestMove(ConnectFour{board, 1}, 1000))
}

func Test_findBestMove_createFork(t *testing.T) {
	var board [7][6]int
	board[2][0] = 1
	board[2][1] = -1
	board[2][2] = 1
	board[3][0] = -1
	board[3][1] = 1
	board[4][0] = -1
	board[3][2] = 1
	board[4][1] = -1
	board[3][3] = 1
	board[3][4] = -1

	assert.Equal(t, ConnectFourAction{4, 1}, findBestMove(ConnectFour{board, 1}, 10000))
}
