package main

import (
	"fmt"
	"will.com/connect4"
	"will.com/mcts"
)

func main() {
	var board [7][6]int
	board[3][0] = 1
	board[3][1] = 1
	board[3][2] = 1
	//board[3][3] = -1
	//board[3][3] = -1
	//board[1][0] = 1k
	//board[2][0] = 1

	var cBoard = connect4.NewConnectFour(board, 1)

	fmt.Println(mcts.NewSearchTree(mcts.RandomRollout{}, float32(1.414), mcts.FlatPolicy{}).FindBestMove(cBoard, 1000))
}
