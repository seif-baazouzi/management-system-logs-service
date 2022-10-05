package handlers

import (
	"fmt"
	"logs-service/src/models"
	"logs-service/src/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateLog(c *fiber.Ctx) error {
	userID := fmt.Sprintf("%s", c.Locals("uuid"))
	workspaceID := c.Params("workspaceID")

	// parse body
	var body models.LogBody

	err := c.BodyParser(&body)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	// check body
	if body.Label == "" {
		return c.JSON(fiber.Map{"label": "Must not be empty"})
	}

	// create log
	logID, err := models.CreateLog(body, workspaceID, userID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success", "logID": logID})
}
