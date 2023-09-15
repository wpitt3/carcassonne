package mcts

type FlatPolicy struct{}

func (_ FlatPolicy) DefinePolicy(state *Node) []float32 {
	var policy []float32
	for i := 0; i < len(state.children); i++ {
		policy = append(policy, 1.0)
	}
	return policy
}
