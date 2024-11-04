package route

import (
	"evermosTest/internal/handler/http/order"
	"evermosTest/internal/handler/http/product"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(productHandler product.Handler, orderHandler order.Handler) *fiber.App {
	router := fiber.New()

	r := router.Group("/")
	r.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("OK")
	})

	apiProduct := r.Group("/product")
	apiProduct.Post("/add", productHandler.CreateNew)
	apiProduct.Get("/list", productHandler.GetList)
	apiProduct.Delete("/:productId", productHandler.Delete)
	apiProduct.Post("/stockIn", productHandler.StockIn)
	apiProduct.Post("/price", productHandler.PriceAdjust)

	router.Post("/checkout", productHandler.Checkout)

	apiOrder := router.Group("/order")
	apiOrder.Get("/list", orderHandler.GetAll)

	return router
}
