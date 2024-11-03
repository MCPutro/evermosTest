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

func (h *handlerImpl) GetProductList(c *fiber.Ctx) error {
	products, err := h.product.GetAllProducts(c.UserContext())
	// error message
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
	return c.Status(fiber.StatusOK).
		JSON(response.Response{
			Code:         fiber.StatusOK,
			Success:      true,
			Message:      "Success",
			TotalRecords: len(products),
			Data:         products,
		})
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
	// error message
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
	body := new(request.Checkout)

	// parse body from request
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
	order, err := h.product.Checkout(c.UserContext(), body)
	// error message
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
			Data:    order,
		})
}

func NewProductHandler(product product.Service) Handler {
	return &handlerImpl{product: product}
}
