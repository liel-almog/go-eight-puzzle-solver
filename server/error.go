package server

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func errorHandler(ctx *fiber.Ctx, err error) error {
	if err != nil {
		code := fiber.StatusInternalServerError
		message := fiber.ErrInternalServerError.Message
		timestamp := time.Now().Format(time.RFC3339)
		path := ctx.Path()

		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			code = fiberErr.Code
			message = fiberErr.Message
		}

		// Handle validator errors
		var fieldErr validator.FieldError
		if errors.As(err, &fieldErr) {
			code = fiber.StatusBadRequest
			message = "Bad Request"
		}

		var validationErr validator.ValidationErrors
		if errors.As(err, &validationErr) {
			code = fiber.StatusBadRequest
			message = "Bad Request"
		}

		// Default error handling
		return ctx.Status(code).JSON(fiber.Map{
			"message":   message,
			"timestamp": timestamp,
			"path":      path,
		})
	}

	return nil
}
