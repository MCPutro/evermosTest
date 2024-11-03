package http

import (
	"evermosTest/internal/handler/request"
	"evermosTest/internal/service/product"
	"evermosTest/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type handlerImpl struct {
	product product.Service
}

func (h *handlerImpl) GetProductList(c *fiber.Ctx) error {
	products, err := h.product.GetAllProducts(c.UserContext())
	// error message
	if err != nil {
		return utils.WriteToResponseBody(c, err, fiber.StatusInternalServerError, err.Error(), nil, 0)
	}

	//success message
	return utils.WriteToResponseBody(c, err, fiber.StatusOK, "Success", nil, len(products))
}

func (h *handlerImpl) CreateNewProduct(c *fiber.Ctx) error {
	body := new(request.Product)

	// parse body message from request
	if err := c.BodyParser(&body); err != nil {
		return utils.WriteToResponseBody(c, err, fiber.StatusBadRequest, "invalid request body", nil, 0)
	}

	// call service
	result, err := h.product.Create(c.UserContext(), body)
	// error message
	if err != nil {
		return utils.WriteToResponseBody(c, err, fiber.StatusInternalServerError, err.Error(), nil, 0)
	}

	//success message
	return utils.WriteToResponseBody(c, err, fiber.StatusCreated, "Success", result, 1)
}

func (h *handlerImpl) Checkout(c *fiber.Ctx) error {
	body := new(request.Checkout)

	// parse body from request
	if err := c.BodyParser(&body); err != nil {
		return utils.WriteToResponseBody(c, err, fiber.StatusBadRequest, "invalid request body", nil, 0)
	}

	// call service
	order, err := h.product.Checkout(c.UserContext(), body)
	// error message
	if err != nil {
		return utils.WriteToResponseBody(c, err, fiber.StatusInternalServerError, err.Error(), nil, 0)
	}

	//success message
	return utils.WriteToResponseBody(c, err, fiber.StatusCreated, "Success", order, 1)
}

func NewProductHandler(product product.Service) Handler {
	return &handlerImpl{product: product}
}
