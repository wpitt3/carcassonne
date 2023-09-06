package main

import "fmt"

func main() {
	var board [7][6]int
	board[0][0] = 1
	board[1][0] = 1
	board[2][0] = 1

	var cBoard = ConnectFour{board, 1}

	fmt.Println(findBestMove(cBoard, 1000))
}
