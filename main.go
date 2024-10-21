package main

import (
	"github.com/all_is_safe/infrastructure/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// Custom CORS configuration

	routers.NewEmergencyRouter(app)
	app.Listen(":3003")
}
