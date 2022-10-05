package handlers

import (
	"fmt"
	"logs-service/src/models"
	"logs-service/src/utils"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

func GetLogs(c *fiber.Ctx) error {
	workspaceID := c.Params("workspaceID")

	// validate month
	month := c.Params("month")

	matched, err := regexp.Match("^[0-9]{4}-[0-9]{1,2}$", []byte(month))

	if err != nil {
		return utils.ServerError(c, err)
	}

	if !matched {
		return c.JSON(fiber.Map{"message": "invalid-month-format"})
	}

	// get logs
	var logs []models.Log

	statingDate := fmt.Sprintf("%s-%d", month, 1)
	endingDate := fmt.Sprintf("%s-%d", month, utils.GetMonthDays(month))

	logs, err = models.GetLogs(workspaceID, statingDate, endingDate)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"logs": logs})
}
