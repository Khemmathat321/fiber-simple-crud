package router

import (
	"github.com/gofiber/fiber/v2"
)

// Registers func
func Registers(app *fiber.App) {
	Order(app)
}

// User func
func Order(router fiber.Router) {
	router = router.Group("/orders")

	// Basic resouce routes
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Index")
	})

	router.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("Post")
	})

	router.Put("/:id", func(c *fiber.Ctx) error {
		return c.SendString("Put" + c.Params("id"))
	})

	router.Delete("/:id", func(c *fiber.Ctx) error {
		return c.SendString("Delete" + c.Params("id"))
	})
}
