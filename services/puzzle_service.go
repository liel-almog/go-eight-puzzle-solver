package services

import (
	"context"
	"sync"

	"github.com/lielalmog/go-be-eight-puzzle-solver/algorithm"
	"github.com/lielalmog/go-be-eight-puzzle-solver/models/dto"
)

type PuzzleService interface {
	GeneratePuzzle(context.Context, *dto.BoardDimensionsDTO) (algorithm.Tiles, error)
	Solve(context.Context, algorithm.Tiles) (algorithm.TilesArray, error)
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

func (p *puzzleServiceImpl) GeneratePuzzle(ctx context.Context, bDimensions *dto.BoardDimensionsDTO) (algorithm.Tiles, error) {
	b, err := algorithm.NewBoard(bDimensions.RowCount, bDimensions.ColumnCount)

	if err != nil {
		return nil, err
	}

	return b.GetTiles(), nil
}

func (p *puzzleServiceImpl) Solve(ctx context.Context, tiles algorithm.Tiles) (algorithm.TilesArray, error) {
	type BfsResult struct {
		solution algorithm.TilesArray
		err      error
	}

	b, err := algorithm.NewBoardFromTiles(tiles)
	if err != nil {
		return nil, err
	}

	bSolver := algorithm.NewBfsSolver(b)
	targetBoard := algorithm.GenerateTargetBoard(b.GetRowCount(), b.GetColumnCount())

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
