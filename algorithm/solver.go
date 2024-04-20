package algorithm

import "errors"

type Tiles = [][]int
type TilesArray = []Tiles

type Solver interface {
	Solve(targetBoard []int) (TilesArray, error)
}

var ErrNoSolution = errors.New("could not find a solution")
