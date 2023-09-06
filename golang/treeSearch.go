package main

import (
	"fmt"
	"math"
	"math/rand"
	// 	"fmt"
)

type Action interface {
}

type Board[A Action] interface {
	Copy() Board[A]
	ValidActions() []A
	PerformMove(A) Board[A]
	IsEndState() bool
	Winner() int
	CurrentPlayer() int
}

type node struct {
	numer             float32
	denom             float32
	board             Board[Action]
	parent            *node
	children          []*node
	isTerminalState   bool
	action            Action
	unexpandedActions []Action
	winner            int
}

func rootNode(board Board[Action]) *node {
	return &node{
		board:             board.Copy(),
		action:            nil,
		unexpandedActions: shuffleActions(board.ValidActions()),
	}
}

func newNode(parentNode *node, action Action) *node {
	board := parentNode.board.PerformMove(action)
	winner := board.Winner()
	isTerminalState := winner != 0 || board.IsEndState()
	newNode := &node{
		board:             board,
		action:            action,
		parent:            parentNode,
		winner:            winner,
		isTerminalState:   isTerminalState,
		unexpandedActions: shuffleActions(board.ValidActions()),
	}
	parentNode.children = append(parentNode.children, newNode)
	return newNode
}

func findBestMove(board Board[Action], turns int) Action {
	rootNode := rootNode(board)
	for i := 0; i < turns; i++ {
		leafNode := selectLeafNode(rootNode)
		result := rolloutGame(leafNode.board)
		var score float32 = 0.0
		if result == 0 {
			score = 0.5
		} else if result != board.CurrentPlayer() {
			score = 1.0
		}
		currentNode := leafNode
		for currentNode != rootNode {
			currentNode.numer += score
			currentNode.denom += 1
			score = 1.0 - score
			currentNode = currentNode.parent
		}
		currentNode.numer += score
		currentNode.denom += 1
	}

	fmt.Println("")
	//fmt.Println(rootNode.children[0].action)
	//fmt.Println(rootNode.children[0].numer)
	//fmt.Println(rootNode.children[0].denom)

	bestMove := rootNode.children[0].action
	bestScore := rootNode.children[0].numer / rootNode.children[0].denom
	for i := 1; i < len(rootNode.children); i++ {
		child := rootNode.children[i]
		//fmt.Println("")
		//fmt.Println(child.action)
		//fmt.Println(child.numer)
		//fmt.Println(child.denom)

		if (child.numer / child.denom) > bestScore {
			bestScore = (child.numer / child.denom)
			bestMove = child.action
		}
	}
	return bestMove
}

func selectLeafNode(rootNode *node) *node {
	currentNode := rootNode
	c := float32(1.414)
	for !currentNode.isTerminalState {

		if len(currentNode.unexpandedActions) > 0 {
			newNode := newNode(currentNode, currentNode.unexpandedActions[0])
			currentNode.unexpandedActions = currentNode.unexpandedActions[1:]
			return newNode
		}
		currentNode = findBestChild(currentNode, c)
	}

	return currentNode
}

func findBestChild(parent *node, c float32) *node {
	logTotalParent := math.Log(float64(parent.denom))
	var maxScore float32 = 0.0
	var bestChild *node = parent.children[0]
	for i := 0; i < len(parent.children); i++ {
		child := parent.children[i]
		score := exploreFunction(child.numer, child.denom, logTotalParent, c)
		if score > maxScore {
			maxScore = score
			bestChild = child
		}
	}
	return bestChild
}

func exploreFunction(wins float32, total float32, logTotalParent float64, c float32) float32 {
	return wins/total + c*float32(math.Sqrt(logTotalParent/float64(total)))
}

func shuffleActions(list []Action) []Action {
	rand := rand.New(rand.NewSource(1))
	for i := len(list); i > 0; i-- {
		ri := rand.Intn(i)
		temp := list[ri]
		list[ri] = list[i-1]
		list[i-1] = temp
	}
	return list
}

func rolloutGame(originalBoard Board[Action]) int {
	rand := rand.New(rand.NewSource(1))
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

	return result
}
