package algorithm

import (
	"errors"

	"github.com/lielalmog/go-be-eight-puzzle-solver/board"
)

type Solver interface {
	Solve(targetBoard []int) (board.TilesArray, error)
}

var (
	ErrNoSolution = errors.New("could not find a solution")
)
