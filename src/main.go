package main

import (
	"fmt"
	"logs-service/src/auth"
	"logs-service/src/db"
	"logs-service/src/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	db.InitDataBase()
	defer db.CloseDataBase()

	app := fiber.New()
	app.Use(cors.New())

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/api/v1/logs/:workspaceID/month/:month", auth.IsWorkspaceOwner, handlers.GetLogs)
	app.Post("/api/v1/logs/:workspaceID", auth.IsWorkspaceOwner, handlers.CreateLog)
	app.Put("/api/v1/logs/:logID", auth.IsLogin, handlers.UpdateLog)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
