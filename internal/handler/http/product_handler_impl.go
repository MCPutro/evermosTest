package http

import (
	"errors"
	"evermosTest/internal/handler/request"
	"evermosTest/internal/service/product"
	"evermosTest/internal/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type handlerImpl struct {
	product product.Service
}

func (h *handlerImpl) GetProductList(c *fiber.Ctx) error {

	// call product service for get all product
	products, err := h.product.GetAllProducts(c.UserContext())

	// return error message
	if err != nil {
		return utils.WriteToResponseBody(c, err, fiber.StatusInternalServerError, err.Error(), nil, 0)
	}

	// return success message
	return utils.WriteToResponseBody(c, err, fiber.StatusOK, "Success", products, len(products))
}

func (h *handlerImpl) CreateNewProduct(c *fiber.Ctx) error {
	body := new(request.Product)

	// parse body message from request
	if err := c.BodyParser(&body); err != nil {
		return utils.WriteToResponseBody(c, err, fiber.StatusBadRequest, "invalid request body", nil, 0)
	}

	// call product service for add new product
	result, err := h.product.Create(c.UserContext(), body)

	// return error message
	if err != nil {
		return utils.WriteToResponseBody(c, err, fiber.StatusInternalServerError, err.Error(), nil, 0)
	}

	// return success message
	return utils.WriteToResponseBody(c, err, fiber.StatusCreated, "Success", result, 1)
}

func (h *handlerImpl) Checkout(c *fiber.Ctx) error {
	body := new(request.Checkout)

	// parse body message from request
	if err := c.BodyParser(&body); err != nil {
		return utils.WriteToResponseBody(c, err, fiber.StatusBadRequest, "invalid request body", nil, 0)
	}

	// Call product service for checkout
	order, err := h.product.Checkout(c.UserContext(), body)

	// return error message
	if err != nil {
		return utils.WriteToResponseBody(c, err, fiber.StatusInternalServerError, err.Error(), nil, 0)
	}

	// return success message
	return utils.WriteToResponseBody(c, err, fiber.StatusCreated, "Success", order, 1)
}

func (h *handlerImpl) DeleteProduct(c *fiber.Ctx) error {
	// get query param input
	sProductId := c.Params("productId", "-1")
	productId, err := strconv.Atoi(sProductId)
	if err != nil {
		return utils.WriteToResponseBody(c, err, fiber.StatusBadRequest, fmt.Sprintf("product id %s is not valid.", sProductId), nil, 0)
	}

	// call product service for delete product
	err = h.product.Delete(c.UserContext(), productId)

	// error handling, return error message
	if err != nil {
		if errors.Is(err, fiber.ErrNotFound) {
			return utils.WriteToResponseBody(c, err, fiber.StatusNotFound, "failed to delete product. "+err.Error(), nil, 0)
		}
		return utils.WriteToResponseBody(c, err, fiber.StatusInternalServerError, "failed to delete user. "+err.Error(), nil, 0)
	}

	// return success message
	return utils.WriteToResponseBody(c, err, fiber.StatusOK, "success", nil, 0)
}

func NewProductHandler(product product.Service) Handler {
	return &handlerImpl{product: product}
}
