package mcts

import (
	"math"
	"math/rand"
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
	return createNode(board)
}

func newNode(parentNode *node, action Action) *node {
	board := parentNode.board.PerformMove(action)
	newNode := createNode(board)
	newNode.parent = parentNode
    newNode.action = action
	parentNode.children = append(parentNode.children, newNode)
	return newNode
}

func createNode(board Board[Action]) *node {
    winner := board.Winner()
    isTerminalState := winner != 0 || board.IsEndState()
    var actions []Action
    if !isTerminalState {
        actions = shuffleActions(board.ValidActions())
    }
    newNode := &node{
        board:             board,
        winner:            winner,
        isTerminalState:   isTerminalState,
        unexpandedActions: actions,
    }
    return newNode
}

func FindBestMove(board Board[Action], turns int) Action {
	rootNode := rootNode(board)
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

func selectLeafNode(rootNode *node) *node {
	currentNode := rootNode
	c := float32(2.414)
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
//
// func printTree(aNode *node, depth int, maxDepth int, minSize int) {
//     sortChildren(aNode.children)
//     if depth < maxDepth {
//         for i := 0; i < len(aNode.children); i++ {
//             child := aNode.children[i]
//             terminal := ""
//             if child.isTerminalState {
//                 terminal = " T"
//             }
//             fmt.Println(strings.Repeat("- ", depth) + fmt.Sprintf("%.3f",child.numer/child.denom) + " " + fmt.Sprint(child.action) + " " + fmt.Sprintf("%.0f", child.denom) + terminal)
//             if child.denom > float32(minSize) {
//                 printTree(child, depth+1, maxDepth, minSize)
//             }
//         }
//     }
// }
//
// func sortChildren(children []*node){
//     sort.Slice(children, func(i, j int) bool {
//         return children[i].action.(ConnectFourAction).index < children[j].action.(ConnectFourAction).index
//     })
// }
