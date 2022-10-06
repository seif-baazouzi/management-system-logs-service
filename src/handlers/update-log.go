package handlers

import (
	"fmt"
	"logs-service/src/models"
	"logs-service/src/utils"

	"github.com/gofiber/fiber/v2"
)

func UpdateLog(c *fiber.Ctx) error {
	userID := fmt.Sprintf("%s", c.Locals("uuid"))
	logID := c.Params("logID")

	// parse body
	var body models.LogBody

	err := c.BodyParser(&body)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	// checks
	if body.Label == "" {
		return c.JSON((fiber.Map{"label": "Must not be empty"}))
	}

	// update log
	err = models.UpdateLog(body, logID, userID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}
