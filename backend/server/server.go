package server

import (
	"context"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const addr = ":5000"

var app *echo.Echo

func Serve() {
	app = echo.New()
	app.HTTPErrorHandler = func(err error, c echo.Context) {}

	app.Use(middleware.Recover())

	setupRouter(app)

	fmt.Println("Server strating on port", addr)

	if err := app.Start(addr); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Shutdown(ctx context.Context) error {
	err := app.Shutdown(ctx)

	if err != nil {
		return err
	}

	return nil
}
