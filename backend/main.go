package main

import (
	"log"

	"github.com/barwareng/sveltekit-golang-supertokens-auth/pkg/config"
	"github.com/barwareng/sveltekit-golang-supertokens-auth/pkg/database"
	"github.com/barwareng/sveltekit-golang-supertokens-auth/pkg/middleware"
	"github.com/barwareng/sveltekit-golang-supertokens-auth/pkg/routes"
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
