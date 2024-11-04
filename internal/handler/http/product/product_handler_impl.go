package product

import (
	"errors"
	"evermosTest/internal/handler/request"
	"evermosTest/internal/service/product"
	"evermosTest/internal/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
)

type handlerImpl struct {
	product product.Service
}

func (h *handlerImpl) GetList(c *fiber.Ctx) error {

	// call product service for get all product
	products, err := h.product.GetAllProducts(c.UserContext())

	// return error message
	if err != nil {
		return utils.WriteToResponseBody(c, err, fiber.StatusInternalServerError, err.Error(), nil, 0)
	}

	// return success message
	return utils.WriteToResponseBody(c, err, fiber.StatusOK, "Success", products, len(products))
}

func (h *handlerImpl) CreateNew(c *fiber.Ctx) error {
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

func (h *handlerImpl) Delete(c *fiber.Ctx) error {
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.WriteToResponseBody(c, err, fiber.StatusNotFound, "failed to delete product. "+err.Error(), nil, 0)
		}
		return utils.WriteToResponseBody(c, err, fiber.StatusInternalServerError, "failed to delete user. "+err.Error(), nil, 0)
	}

	// return success message
	return utils.WriteToResponseBody(c, err, fiber.StatusOK, "Success", nil, 0)
}

func (h *handlerImpl) StockIn(c *fiber.Ctx) error {
	body := new(request.ProductStockIn)

	// parse body message from request
	if err := c.BodyParser(&body); err != nil {
		return utils.WriteToResponseBody(c, err, fiber.StatusBadRequest, "invalid request body", nil, 0)
	}

	if body.AddStock == 0 {
		return utils.WriteToResponseBody(c, errors.New("value of add_stock cannot be zero"), fiber.StatusBadRequest, "invalid request body", nil, 0)
	}

	isDeduct := false
	if body.AddStock < 0 {
		isDeduct = true
		body.AddStock = body.AddStock * -1
	}

	fmt.Println("isDeduct :", isDeduct)
	fmt.Println("body:", body)

	// call product service for data product adjustment
	adjustment, err := h.product.PriceAndStockAdjustment(c.UserContext(), body.ProductId, isDeduct, body.AddStock, -1)

	// return error message
	if err != nil {
		return utils.WriteToResponseBody(c, err, fiber.StatusInternalServerError, err.Error(), nil, 0)
	}

	// return success message
	return utils.WriteToResponseBody(c, err, fiber.StatusCreated, "Stock updated successfully", adjustment, 1)
}

func (h *handlerImpl) PriceAdjust(c *fiber.Ctx) error {
	body := new(request.ProductPriceAdjust)

	// parse body message from request
	if err := c.BodyParser(&body); err != nil {
		return utils.WriteToResponseBody(c, err, fiber.StatusBadRequest, "invalid request body", nil, 0)
	}

	if body.NewPrice < 0 {
		return utils.WriteToResponseBody(c, errors.New("value of new_price cannot be less then zero"), fiber.StatusBadRequest, "invalid request body", nil, 0)
	}

	// call product service for data product adjustment
	adjustment, err := h.product.PriceAndStockAdjustment(c.UserContext(), body.ProductId, false, -1, body.NewPrice)

	// return error message
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.WriteToResponseBody(c, err, fiber.StatusNotFound, "failed to update product. "+err.Error(), nil, 0)
		}
		return utils.WriteToResponseBody(c, err, fiber.StatusInternalServerError, err.Error(), nil, 0)
	}

	// return success message
	return utils.WriteToResponseBody(c, err, fiber.StatusCreated, "Price updated successfully", adjustment, 1)
}

func NewProductHandler(product product.Service) Handler {
	return &handlerImpl{product: product}
}
