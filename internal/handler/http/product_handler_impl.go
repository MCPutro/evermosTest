package http

import (
	"evermosTest/internal/handler/request"
	"evermosTest/internal/handler/response"
	"evermosTest/internal/service/product"
	"github.com/gofiber/fiber/v2"
)

type handlerImpl struct {
	product product.Service
}

func (h *handlerImpl) CreateNewProduct(c *fiber.Ctx) error {
	body := new(request.Product)

	// parse body message from request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.Response{
				Code:    fiber.StatusBadRequest,
				Message: "invalid request body",
				Success: false,
				Data:    nil,
			})
	}

	// call service
	result, err := h.product.Create(c.UserContext(), body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(response.Response{
				Code:    fiber.StatusInternalServerError,
				Message: err.Error(),
				Success: false,
				Data:    nil,
			})
	}

	//success message
	return c.Status(fiber.StatusCreated).
		JSON(response.Response{
			Code:    fiber.StatusCreated,
			Success: true,
			Message: "Created",
			Data:    result,
		})
}

func (h *handlerImpl) Checkout(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func NewProductHandler(product product.Service) Handler {
	return &handlerImpl{product: product}
}
