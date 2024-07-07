package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lielalmog/go-eight-puzzle-solver/routes"
)

func setupRouter(app *echo.Echo) {
	app.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Hello World!",
		})
	})

	api := app.Group("/api")
	routes.NewPuzzleRouter(api)
}
