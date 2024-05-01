package main

import (
	"log"

	"github.com/acme-corp/pkg/config"
	"github.com/acme-corp/pkg/database"
	"github.com/acme-corp/pkg/middleware"
	"github.com/acme-corp/pkg/routes"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	database.ConnectDb()
	config.SupertokensInit()
	app := fiber.New()
	middleware.FiberMiddleware(app)
	routes.InitRoutes(app)
	if err := app.Listen(":3000"); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
