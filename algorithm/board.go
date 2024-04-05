package algorithm

import (
	"github.com/lielalmog/go-be-eight-puzzle-solver/errors/apperrors"
)

const BoardBlackValue = -1

type Board struct {
	rowCount    int
	columnCount int
	tiles       [][]int
}

func NewBoard(rowCount, columnCount int) (*Board, error) {
	if rowCount < 1 {
		return nil, apperrors.ErrRowCountSmallerThanOne
	}

	if columnCount < 1 {
		return nil, apperrors.ErrColumnCountSmallerThanOne
	}

	board := &Board{
		rowCount:    rowCount,
		columnCount: columnCount,
	}

	return board, nil
}
