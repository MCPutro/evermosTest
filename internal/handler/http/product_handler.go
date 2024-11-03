package http

import "github.com/gofiber/fiber/v2"

type Handler interface {
	CreateNewProduct(c *fiber.Ctx) error
	Checkout(c *fiber.Ctx) error
	GetProductList(c *fiber.Ctx) error
	DeleteProduct(c *fiber.Ctx) error
}
