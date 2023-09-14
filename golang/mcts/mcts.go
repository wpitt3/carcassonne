package mcts

import (
	"math"
	"math/rand"
)

const C = float32(1.414)

type Action interface {
}

type State[A Action] interface {
	Copy() State[A]
	ValidActions() []A
	PerformMove(A) State[A]
	IsEndState() bool
	Winner() int
	CurrentPlayer() int
}

func FindBestMove(board State[Action], turns int) Action {
	rootNode := createNode(board)
	for i := 0; i < turns; i++ {
		leafNode := selectLeafNode(rootNode)
		result := rolloutGame(leafNode.board)
		var score float32 = 0.0
		if result == 0 {
			score = 0.5
		} else if result == board.CurrentPlayer() {
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

	bestMove := rootNode.children[0].action
	bestScore := rootNode.children[0].numer / rootNode.children[0].denom
	for i := 1; i < len(rootNode.children); i++ {
		child := rootNode.children[i]
		if (child.numer / child.denom) > bestScore {
			bestScore = (child.numer / child.denom)
			bestMove = child.action
		}
	}
	return bestMove
}

type node struct {
	numer             float32
	denom             float32
	board             State[Action]
	parent            *node
	children          []*node
	isTerminal        bool
	action            Action
	unexpandedActions []Action
	winner            int
}

func createNode(board State[Action]) *node {
	winner := board.Winner()
	isTerminalState := winner != 0 || board.IsEndState()
	var actions []Action
	if !isTerminalState {
		actions = shuffleActions(board.ValidActions())
	}
	newNode := &node{
		board:             board,
		winner:            winner,
		isTerminal:        isTerminalState,
		unexpandedActions: actions,
	}
	return newNode
}

func newNode(parentNode *node, action Action) *node {
	board := parentNode.board.PerformMove(action)
	newNode := createNode(board)
	newNode.parent = parentNode
	newNode.action = action
	parentNode.children = append(parentNode.children, newNode)
	return newNode
}

func selectLeafNode(rootNode *node) *node {
	currentNode := rootNode
	for !currentNode.isTerminal {
		if len(currentNode.unexpandedActions) > 0 {
			newNode := newNode(currentNode, currentNode.unexpandedActions[0])
			currentNode.unexpandedActions = currentNode.unexpandedActions[1:]
			return newNode
		}
		currentNode = findBestChild(currentNode)
	}

	return currentNode
}

func findBestChild(parent *node) *node {
	logTotalParent := math.Log(float64(parent.denom))
	var maxScore float32 = 0.0
	var bestChild = parent.children[0]
	for i := 0; i < len(parent.children); i++ {
		child := parent.children[i]
		score := exploreFunction(child.numer, child.denom, logTotalParent)
		if score > maxScore {
			maxScore = score
			bestChild = child
		}
	}
	return bestChild
}

func exploreFunction(wins float32, total float32, logTotalParent float64) float32 {
	return wins/total + C*float32(math.Sqrt(logTotalParent/float64(total)))
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

func rolloutGame(originalBoard State[Action]) int {
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
