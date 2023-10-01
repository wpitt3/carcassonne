package mcts

type RolloutEngine interface {
	Rollout(State[Action]) float32
}

type PolicyEngine interface {
	DefinePolicy(State[Action], []Action) []float32
}

type Action interface {
}

type State[A Action] interface {
	Copy() State[A]
	ValidActions() []A
	PerformMove(A) State[A]
	IsEndState() bool
	Winner() int
	PrintState()
}
