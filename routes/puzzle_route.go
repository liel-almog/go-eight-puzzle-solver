package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/lielalmog/go-be-eight-puzzle-solver/controllers"
)

func NewPuzzleRouter(router *echo.Group) {
	group := router.Group("/puzzle")

	controller := controllers.GetPuzzleController()
	group.POST("/generate", controller.GeneratePuzzle)
	group.POST("/bfs/solve", controller.BfsSolve)
	group.POST("/dfs/solve", controller.DfsSolve)
	group.POST("/astar/solve", controller.AStarSolve)
}
