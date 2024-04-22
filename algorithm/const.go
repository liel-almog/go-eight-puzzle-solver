package algorithm

import "errors"

const BoardBlankValue = -1

var (
	ErrNoSolution       = errors.New("could not find a solution")
	ErrEmptyTiles       = errors.New("the tiles array is empty")
	ErrTilesNotSameSize = errors.New("not same size tile")
)
