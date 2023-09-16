package main

import (
	"fmt"
	"will.com/connect4"
	"will.com/mcts"
)

func main() {
	var treeA = mcts.NewSearchTree(mcts.RandomRollout{}, float32(1.414), mcts.FlatPolicy{})
	var treeB = mcts.NewSearchTree(mcts.RandomRollout{}, float32(1.414), mcts.FlatPolicy{})

	games := 100
	startPlayer := 1
	for i := 0; i < games; i++ {
		var cBoard = connect4.NewConnectFour([7][6]int{}, startPlayer)
		for !cBoard.IsEndState() && cBoard.Winner() == 0 {
			var move mcts.Action
			if cBoard.CurrentPlayer() == 1 {
				move = treeA.FindBestMoveByTime(cBoard, 300)
			} else {
				move = treeB.FindBestMoveByTime(cBoard, 300)
			}
			cBoard = cBoard.PerformMove(move).(connect4.ConnectFour)
		}
		startPlayer *= -1
		fmt.Println(cBoard.Winner())
		cBoard.PrintState()
	}
}
