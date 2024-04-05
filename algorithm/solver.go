package algorithm

type Solver interface {
	Solve() (*Board, error)
}
