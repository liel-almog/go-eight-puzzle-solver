package algorithm

import "github.com/lielalmog/go-be-eight-puzzle-solver/board"

type AStarSolver struct{}

func NewAStarSolver() Solver {
	return &AStarSolver{}
}

func (aStar *AStarSolver) Solve(targerBoard []int) (board.TilesArray, error) {
	return nil, nil
}
