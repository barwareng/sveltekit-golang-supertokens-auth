package routes

import (
	"github.com/barwareng/sveltekit-golang-supertokens-auth/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func InitRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server is running")
	})
	protectedRoutes := app.Group("/api", adaptor.HTTPMiddleware(middleware.VerifySession))
	protectedRoutes.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("I am protected")
	})

}
