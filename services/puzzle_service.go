package services

import (
	"context"
	"sync"

	"github.com/lielalmog/go-be-eight-puzzle-solver/algorithm"
	"github.com/lielalmog/go-be-eight-puzzle-solver/board"
	"github.com/lielalmog/go-be-eight-puzzle-solver/models/dto"
)

type PuzzleService interface {
	GeneratePuzzle(context.Context, *dto.BoardDimensionsDTO) (board.Tiles, error)
	Solve(context.Context, board.Tiles) (board.TilesArray, error)
}

type puzzleServiceImpl struct{}

var (
	initPuzzleService sync.Once
	puzzleService     *puzzleServiceImpl
)

func newPuzzleService() *puzzleServiceImpl {
	return &puzzleServiceImpl{}
}

func GetPuzzleService() PuzzleService {
	initPuzzleService.Do(func() {
		puzzleService = newPuzzleService()
	})

	return puzzleService
}

func (p *puzzleServiceImpl) GeneratePuzzle(ctx context.Context, bDimensions *dto.BoardDimensionsDTO) (board.Tiles, error) {
	b, err := board.NewBoard(bDimensions.RowCount, bDimensions.ColumnCount)

	if err != nil {
		return nil, err
	}

	return b.GetTiles(), nil
}

func (p *puzzleServiceImpl) Solve(ctx context.Context, tiles board.Tiles) (board.TilesArray, error) {
	type BfsResult struct {
		solution board.TilesArray
		err      error
	}

	b, err := board.NewBoardFromTiles(tiles)
	if err != nil {
		return nil, err
	}

	bSolver := algorithm.NewBfsSolver(b)
	targetBoard := board.GenerateTargetBoard(b.GetRowCount(), b.GetColumnCount())

	ch := make(chan BfsResult)

	go func() {
		solution, err := bSolver.Solve(targetBoard)

		ch <- BfsResult{
			solution: solution,
			err:      err,
		}
	}()

	select {
	case res := <-ch:
		if res.err != nil {
			return nil, err
		}

		return res.solution, nil

	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
