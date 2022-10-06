package handlers

import (
	"fmt"
	"logs-service/src/models"
	"logs-service/src/utils"

	"github.com/gofiber/fiber/v2"
)

func DeleteLog(c *fiber.Ctx) error {
	userID := fmt.Sprintf("%s", c.Locals("uuid"))
	logID := c.Params("logID")

	err := models.DeleteLog(logID, userID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}
