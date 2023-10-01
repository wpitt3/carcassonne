package mcts

import "math/rand"

type RandomRollout struct{}

func (_ RandomRollout) Rollout(originalBoard State[Action]) float32 {
	board := originalBoard.Copy()
	result := board.Winner()
	done := result != 0 || board.IsEndState()
	for !done {
		moves := board.ValidActions()
		move := moves[rand.Intn(len(moves))]
		board = board.PerformMove(move)
		result = board.Winner()
		done = result != 0 || board.IsEndState()
	}
	if result == 0 {
		return 0.5
	} else if result == -1 {
		return 1.0
	}

	return 0.0
}
