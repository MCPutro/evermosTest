package order

import "github.com/gofiber/fiber/v2"

type Handler interface {
	GetAll(c *fiber.Ctx) error
}
