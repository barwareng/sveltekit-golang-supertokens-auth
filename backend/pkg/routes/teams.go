package routes

import (
	"github.com/acme-corp/app/controllers"
	"github.com/acme-corp/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func teamRoutes(app *fiber.App) {
	route := app.Group("/teams", adaptor.HTTPMiddleware(middleware.VerifySession))
	route.Post("/create", controllers.AddTeam)
	route.Get("/", controllers.GetTeams)
}
