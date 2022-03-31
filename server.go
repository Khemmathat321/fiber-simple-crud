package main

import (
	router "simple-crud/app/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.Registers(app)

	app.Listen(":80")
}
