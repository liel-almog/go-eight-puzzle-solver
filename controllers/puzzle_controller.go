package controllers

import (
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/lielalmog/go-be-eight-puzzle-solver/configs"
	"github.com/lielalmog/go-be-eight-puzzle-solver/models/dto"
	"github.com/lielalmog/go-be-eight-puzzle-solver/services"
)

type PuzzleController interface {
	GeneratePuzzle(c *fiber.Ctx) error
	Solve(c *fiber.Ctx) error
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

func (p *puzzleControllerImpl) GeneratePuzzle(c *fiber.Ctx) error {
	bDimensions := new(dto.BoardDimensionsDTO)

	if err := c.BodyParser(bDimensions); err != nil {
		return fiber.ErrBadRequest
	}

	if err := configs.GetValidator().Struct(bDimensions); err != nil {
		return fiber.ErrBadRequest
	}

	tiles, err := p.puzzleService.GeneratePuzzle(c.Context(), bDimensions)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(tiles)
}

func (p *puzzleControllerImpl) Solve(c *fiber.Ctx) error {
	tiles := new(dto.TilesDTO)

	if err := c.BodyParser(tiles); err != nil {
		return fiber.ErrBadRequest
	}

	if err := configs.GetValidator().Struct(tiles); err != nil {
		return fiber.ErrBadRequest
	}

	solution, err := p.puzzleService.Solve(c.Context(), tiles.Tiles)
	if err != nil {
		return err
	}

	return c.JSON(solution)
}
