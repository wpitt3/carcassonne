package mcts

type TicTacToe struct {
	board  [3][3]int
	player int
}

type TicTacToeAction struct {
	x      int
	y      int
	player int
}

func (board TicTacToe) Copy() State[Action] {
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
				return false
			}
		}
	}
	return true
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

func (board TicTacToe) PerformMove(action Action) State[Action] {
	ticTacToeAction := action.(TicTacToeAction)
	board.board[ticTacToeAction.x][ticTacToeAction.y] = ticTacToeAction.player
	board.player = board.player * -1
	return board
}
