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
	apiProduct.Post("/add", handler.CreateNew)
	apiProduct.Get("/list", handler.GetList)
	apiProduct.Delete("/:productId", handler.Delete)
	apiProduct.Post("/stockIn", handler.StockIn)
	apiProduct.Post("/price", handler.PriceAdjust)

	router.Post("/checkout", handler.Checkout)

	return router
}
