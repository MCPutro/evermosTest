package http

import "github.com/gofiber/fiber/v2"

type Handler interface {
	CreateNew(c *fiber.Ctx) error
	Checkout(c *fiber.Ctx) error
	GetList(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	StockIn(c *fiber.Ctx) error
	PriceAdjust(c *fiber.Ctx) error
}
