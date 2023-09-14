package mcts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


type TicTacToe struct {
	board  [3][3]int
	player int
}

type TicTacToeAction struct {
	x  int
	y  int
	player int
}

func (board TicTacToe) Copy() Board[Action] {
	var newBoard TicTacToe
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			newBoard.board[i][j] = board.board[i][j]
		}
	}
	newBoard.player = board.player
	return newBoard
}

func (board TicTacToe) ValidActions() []Action {
	validMoves := make([]Action, 0)
	for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board.board[i][j] == 0 {
			    validMoves = append(validMoves, TicTacToeAction{i, j, board.player})
			}
		}
	}
	return validMoves
}

func (board TicTacToe) IsEndState() bool {
    for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
            if board.board[i][j] == 0 {
                return false;
            }
		}
    }
	return true;
}

func (board TicTacToe) Winner() int {
	for i := 0; i < 3; i++ {
	    var sum = 0
        for j := 0; j < 3; j++ {
            sum += board.board[i][j]
        }
        if abs(sum) == 3 {
            return sum / 3
        }
    }
    for i := 0; i < 3; i++ {
        var sum = 0
        for j := 0; j < 3; j++ {
            sum += board.board[j][i]
        }
        if abs(sum) == 3 {
            return sum / 3
        }
    }
    var sum = board.board[0][0] + board.board[1][1] + board.board[2][2]
    if abs(sum) == 3 {
        return sum / 3
    }
    sum = board.board[2][0] + board.board[1][1] + board.board[0][2]
    if abs(sum) == 3 {
        return sum / 3
    }
    return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (board TicTacToe) CurrentPlayer() int {
	return board.player
}

func (board TicTacToe) PerformMove(action Action) Board[Action] {
	ticTacToeAction := action.(TicTacToeAction)
	board.board[ticTacToeAction.x][ticTacToeAction.y] = ticTacToeAction.player
    board.player = board.player * -1
	return board
}

func Test_rootnode(t *testing.T) {
	var board Board[Action] = TicTacToe{[3][3]int{}, 1}
	board = board.PerformMove(TicTacToeAction{0, 0, 1}).(TicTacToe)
	rootNode := rootNode(board)

	assert.Equal(t, nil, rootNode.action)
	assert.Equal(t, 1, rootNode.board.(TicTacToe).board[0][0])
	assert.Equal(t, -1, rootNode.board.(TicTacToe).player)
	assert.Equal(t, 8, len(rootNode.unexpandedActions))
}

func Test_newNode(t *testing.T) {
	var board Board[Action] = TicTacToe{[3][3]int{}, 1}
	rootNode := rootNode(board)
	newNode := newNode(rootNode, TicTacToeAction{0, 0, 1})

	assert.Equal(t, TicTacToeAction{0, 0, 1}, newNode.action)
	assert.Equal(t, 1, newNode.board.(TicTacToe).board[0][0])
	assert.Equal(t, -1, newNode.board.(TicTacToe).player)
	assert.Equal(t, rootNode, newNode.parent)
	assert.Equal(t, 0, newNode.winner)
	assert.Equal(t, false, newNode.isTerminalState)
	assert.Equal(t, 9, len(rootNode.unexpandedActions))
	assert.Equal(t, 8, len(newNode.unexpandedActions))
	assert.Equal(t, rootNode.children[0], newNode)
}

func Test_newNodeHasWinner(t *testing.T) {
	var board Board[Action] = TicTacToe{[3][3]int{}, 1}
	board = board.PerformMove(TicTacToeAction{0, 0, 1})
	board = board.PerformMove(TicTacToeAction{1, 0, 1})
	rootNode := rootNode(board)
	newNode := newNode(rootNode, TicTacToeAction{2, 0, 1})

	assert.Equal(t, TicTacToeAction{2, 0, 1}, newNode.action)
	assert.Equal(t, 1, newNode.winner)
	assert.Equal(t, true, newNode.isTerminalState)
}

func Test_newNodeIsFullWithNoWinner(t *testing.T) {
	var board [3][3]int
    board[0][0] = 1
    board[0][1] = 1
    board[0][2] = -1

    board[1][0] = -1
    board[1][1] = -1
    board[1][2] = 1

    board[2][0] = 1
    board[2][1] = -1
    board[2][2] = 1

	rootNode := rootNode(TicTacToe{board, 1})
	nodeA := newNode(rootNode, TicTacToeAction{0, 0, 1})
	assert.Equal(t, 0, nodeA.winner)
	assert.Equal(t, true, nodeA.isTerminalState)
}

func Test_findBestChild(t *testing.T) {
	var board Board[Action] = TicTacToe{[3][3]int{}, 1}
	rootNode := rootNode(board)
	nodeA := newNode(rootNode, TicTacToeAction{1, 0, 1})
	nodeB := newNode(rootNode, TicTacToeAction{2, 0, 1})

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
	var board Board[Action] = TicTacToe{[3][3]int{}, 1}
	rootNode := rootNode(board)
	action1 := rootNode.unexpandedActions[0]

	node := selectLeafNode(rootNode)
	assert.Equal(t, node.action, action1)
}

//
//func Test_selectLeafNode_expandChild(t *testing.T) {
//	var board TicTacToe
//	rootNode := rootNode(board)
//	rootNode.unexpandedActions = []TicTacToeAction
//	nodeA := newNode(rootNode, 1)
//	action1 := nodeA.unexpandedActions[0]
//
//	node := selectLeafNode(rootNode)
//	assert.Equal(t, node.action, action1)
//}
//
//func Test_selectLeafNode_dontExpandTerminal(t *testing.T) {
//	var board TicTacToe
//	rootNode := rootNode(board)
//	rootNode.unexpandedActions = []TicTacToeAction
//	nodeA := newNode(rootNode, 1)
//	nodeA.isTerminalState = true
//
//	node := selectLeafNode(rootNode)
//	assert.Equal(t, node, nodeA)
//}