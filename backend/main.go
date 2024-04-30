package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New()
	if err := app.Listen(":3000"); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
