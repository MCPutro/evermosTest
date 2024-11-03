package route

import (
	"evermosTest/internal/handler/http"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(handler http.Handler) *fiber.App {
	router := fiber.New()

	r := router.Group("/")
	r.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("OK")
	})

	apiProduct := r.Group("/product")
	apiProduct.Post("/add", handler.CreateNewProduct)
	apiProduct.Get("/list", handler.GetProductList)
	apiProduct.Delete("/:productId", handler.DeleteProduct)

	router.Post("/checkout", handler.Checkout)

	return router
}
