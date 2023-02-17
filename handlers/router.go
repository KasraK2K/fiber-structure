package handlers

import (
	"app/handlers/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Routes(app *fiber.App) {
	app.Get("/_health", health)
	app.Get("/_metrics", monitor.New(monitor.Config{Title: "Default Metrics Page"}))

	api := app.Group("/api")
	v1 := api.Group("/v1")
	users.Routes(v1)
}

func health(c *fiber.Ctx) error {
	return c.SendString("Everything works fine")
}
