package server

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func errorHandler(ctx echo.Context, err error) error {
	if err != nil {
		code := http.StatusInternalServerError
		message := echo.ErrInternalServerError.Message
		timestamp := time.Now().Format(time.RFC3339)
		path := ctx.Path()

		var echoErr *echo.HTTPError
		if errors.As(err, &echoErr) {
			code = echoErr.Code
			message = echoErr.Message
		}

		// Handle validator errors
		var fieldErr validator.FieldError
		if errors.As(err, &fieldErr) {
			code = http.StatusBadRequest
			message = "Bad Request"
		}

		var validationErr validator.ValidationErrors
		if errors.As(err, &validationErr) {
			code = http.StatusBadRequest
			message = "Bad Request"
		}

		// Default error handling
		return ctx.JSON(code, echo.Map{
			"message":   message,
			"timestamp": timestamp,
			"path":      path,
		})
	}

	return nil
}
