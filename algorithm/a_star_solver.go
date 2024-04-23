package algorithm

import "github.com/lielalmog/go-be-eight-puzzle-solver/board"

type AStarSolver struct {
	InitialBoard *board.Board
}

func NewAStarSolver(b *board.Board) Solver {
	return &AStarSolver{
		InitialBoard: b,
	}
}

func (aStar *AStarSolver) Solve(targerBoard []int) (board.TilesArray, error) {
	return nil, nil
}
