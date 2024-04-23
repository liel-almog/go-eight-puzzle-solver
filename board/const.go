package board

import "errors"

const BoardBlankValue = -1

var (
	ErrEmptyTiles       = errors.New("the tiles array is empty")
	ErrTilesNotSameSize = errors.New("not same size tile")
)
