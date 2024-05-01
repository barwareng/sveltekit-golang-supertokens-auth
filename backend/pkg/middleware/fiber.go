package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/supertokens/supertokens-golang/supertokens"
)

// FiberMiddleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/api/middleware
func FiberMiddleware(app *fiber.App) {
	app.Use(
		cors.New(cors.Config{
			AllowOrigins:     os.Getenv("CORS_ALLOWED_ORIGINS"),
			AllowHeaders:     "Origin, Content-Type, Accept, X-Team, " + strings.Join(supertokens.GetAllCORSHeaders(), ", "),
			AllowCredentials: true,
		}),
		recover.New(),
		adaptor.HTTPMiddleware(supertokens.Middleware),
	)
}
