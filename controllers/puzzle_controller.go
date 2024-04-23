package controllers

import (
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/lielalmog/go-be-eight-puzzle-solver/configs"
	"github.com/lielalmog/go-be-eight-puzzle-solver/models/dto"
	"github.com/lielalmog/go-be-eight-puzzle-solver/services"
)

type PuzzleController interface {
	GeneratePuzzle(c echo.Context) error
	BfsSolve(c echo.Context) error
	DfsSolve(c echo.Context) error
	AStarSolve(c echo.Context) error
}

type puzzleControllerImpl struct {
	puzzleService services.PuzzleService
}

var (
	initPuzzleController sync.Once
	puzzleController     *puzzleControllerImpl
)

func newPuzzleController() *puzzleControllerImpl {
	return &puzzleControllerImpl{
		puzzleService: services.GetPuzzleService(),
	}
}

func GetPuzzleController() PuzzleController {
	initPuzzleController.Do(func() {
		puzzleController = newPuzzleController()
	})

	return puzzleController
}

func (p *puzzleControllerImpl) GeneratePuzzle(c echo.Context) error {
	bDimensions := new(dto.BoardDimensionsDTO)

	if err := c.Bind(bDimensions); err != nil {
		return echo.ErrBadRequest
	}

	if err := configs.GetValidator().Struct(bDimensions); err != nil {
		return echo.ErrBadRequest
	}

	tiles, err := p.puzzleService.GeneratePuzzle(c.Request().Context(), bDimensions)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, tiles)
}

func (p *puzzleControllerImpl) BfsSolve(c echo.Context) error {
	tiles := new(dto.TilesDTO)

	if err := c.Bind(tiles); err != nil {
		return echo.ErrBadRequest
	}

	if err := configs.GetValidator().Struct(tiles); err != nil {
		return echo.ErrBadRequest
	}

	solution, err := p.puzzleService.BfsSolve(c.Request().Context(), tiles.Tiles)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, solution)
}

func (p *puzzleControllerImpl) DfsSolve(c echo.Context) error {
	tiles := new(dto.TilesDTO)

	if err := c.Bind(tiles); err != nil {
		return echo.ErrBadRequest
	}

	if err := configs.GetValidator().Struct(tiles); err != nil {
		return echo.ErrBadRequest
	}

	solution, err := p.puzzleService.DfsSolve(c.Request().Context(), tiles.Tiles)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, solution)
}

func (p *puzzleControllerImpl) AStarSolve(c echo.Context) error {
	tiles := new(dto.TilesDTO)

	if err := c.Bind(tiles); err != nil {
		return echo.ErrBadRequest
	}

	if err := configs.GetValidator().Struct(tiles); err != nil {
		return echo.ErrBadRequest
	}

	solution, err := p.puzzleService.AStarSolve(c.Request().Context(), tiles.Tiles)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, solution)
}
