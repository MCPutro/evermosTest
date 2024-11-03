package route

import (
	"evermosTest/internal/handler/http"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(handler http.Handler) *fiber.App {
	router := fiber.New()

	router.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	router.Post("/addProduct", handler.CreateNewProduct)

	return router
}
