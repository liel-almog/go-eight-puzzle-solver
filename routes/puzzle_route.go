package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lielalmog/go-be-eight-puzzle-solver/controllers"
)

func NewPuzzleRouter(router fiber.Router) {
	group := router.Group("/puzzle")

	controller := controllers.GetPuzzleController()
	group.Post("/generate", controller.GeneratePuzzle)
	group.Post("/bfs/solve", controller.BfsSolve)
	group.Post("/dfs/solve", controller.DfsSolve)
	group.Post("/astar/solve", controller.AStarSolve)
}
