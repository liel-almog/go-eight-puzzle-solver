package services

import (
	"sync"

	"github.com/lielalmog/go-be-eight-puzzle-solver/algorithm"
	"github.com/lielalmog/go-be-eight-puzzle-solver/models/dto"
)

type PuzzleService interface {
	GeneratePuzzle(*dto.BoardDimensionsDTO) (algorithm.Tiles, error)
	Solve(algorithm.Tiles) (algorithm.TilesArray, error)
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

func (p *puzzleServiceImpl) GeneratePuzzle(bDimensions *dto.BoardDimensionsDTO) (algorithm.Tiles, error) {
	b, err := algorithm.NewBoard(bDimensions.RowCount, bDimensions.ColumnCount)

	if err != nil {
		return nil, err
	}

	return b.GetTiles(), nil
}

func (p *puzzleServiceImpl) Solve(tiles algorithm.Tiles) (algorithm.TilesArray, error) {
	return nil, nil
}
