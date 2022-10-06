package handlers

import (
	"logs-service/src/models"
	"logs-service/src/utils"

	"github.com/gofiber/fiber/v2"
)

func DeleteWorkspaceLogs(c *fiber.Ctx) error {
	workspaceID := c.Params("workspaceID")

	err := models.DeleteWorkspaceLogs(workspaceID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}
