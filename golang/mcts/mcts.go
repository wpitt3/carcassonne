package mcts

import (
	"math"
	"math/rand"
)

type RolloutEngine interface {
	Rollout(State[Action], int) float32
}

type PolicyEngine interface {
	DefinePolicy(*Node) []float32
}

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

type Node struct {
	numer             float32
	denom             float32
	policyScore       float32
	board             State[Action]
	parent            *Node
	children          []*Node
	isTerminal        bool
	action            Action
	unexpandedActions []Action
	winner            int
}

type SearchTree struct {
	rolloutEngine RolloutEngine
	c             float32
	policyEngine  PolicyEngine
}

func NewSearchTree(rolloutEngine RolloutEngine, c float32, policyEngine PolicyEngine) SearchTree {
	return SearchTree{rolloutEngine, c, policyEngine}
}

func (searchTree SearchTree) FindBestMove(board State[Action], turns int) Action {
	rootNode := createNode(board)
	for i := 0; i < turns; i++ {
		leafNode := searchTree.selectLeafNode(rootNode)
		var score = searchTree.rolloutEngine.Rollout(leafNode.board, board.CurrentPlayer())
		backpropagation(leafNode, rootNode, score)
	}
	return calculateBestAction(rootNode)
}

func backpropagation(leafNode *Node, rootNode *Node, score float32) {
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

func calculateBestAction(rootNode *Node) Action {
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

func createNode(board State[Action]) *Node {
	winner := board.Winner()
	isTerminalState := winner != 0 || board.IsEndState()
	var actions []Action
	if !isTerminalState {
		actions = shuffleActions(board.ValidActions())
	}
	newNode := &Node{
		board:             board,
		winner:            winner,
		isTerminal:        isTerminalState,
		unexpandedActions: actions,
	}
	return newNode
}

func newNode(parentNode *Node, action Action) *Node {
	board := parentNode.board.PerformMove(action)
	newNode := createNode(board)
	newNode.parent = parentNode
	newNode.action = action
	parentNode.children = append(parentNode.children, newNode)
	return newNode
}

func (searchTree SearchTree) selectLeafNode(rootNode *Node) *Node {
	currentNode := rootNode
	for !currentNode.isTerminal {
		if len(currentNode.unexpandedActions) > 0 {
			newNode := newNode(currentNode, currentNode.unexpandedActions[0])
			currentNode.unexpandedActions = currentNode.unexpandedActions[1:]
			return newNode
		}
		if currentNode.children[0].policyScore == 0.0 {
			policy := searchTree.policyEngine.DefinePolicy(currentNode)
			if len(policy) != len(currentNode.children) {
				panic("Policy does not match children")
			}
			for i := 0; i < len(currentNode.children); i++ {
				currentNode.children[i].policyScore = policy[i]
			}
		}
		currentNode = searchTree.findBestChild(currentNode)
	}
	return currentNode
}

func (searchTree SearchTree) findBestChild(parent *Node) *Node {
	logTotalParent := math.Log(float64(parent.denom))
	var maxScore float32 = 0.0
	var bestChild = parent.children[0]
	for i := 0; i < len(parent.children); i++ {
		child := parent.children[i]
		score := exploreFunction(child.numer, child.denom, logTotalParent, searchTree.c)
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
