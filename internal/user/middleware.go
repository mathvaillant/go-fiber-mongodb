package user

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

// If user does not exist, do not allow one to access the API.
func (handler *UserHandler) checkIfUserExistsMiddleware(ctx *fiber.Ctx) error {
	// Create a new customized context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Fetch parameter.
	targetedUserID := ctx.Params("userID")

	if targetedUserID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "User not found!",
		})
	}

	// Check if user exists.
	searchedUser, err := handler.userService.GetUser(customContext, targetedUserID)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if searchedUser == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "User not found!",
		})
	}

	return ctx.Next()
}
