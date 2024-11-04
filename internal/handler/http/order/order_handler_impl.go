package order

import (
	"evermosTest/internal/service/order"
	"evermosTest/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type handlerImpl struct {
	order order.Service
}

func (h *handlerImpl) GetAll(c *fiber.Ctx) error {
	orderList, err := h.order.GetAllOrder(c.UserContext())
	if err != nil {
		return utils.WriteToResponseBody(c, err, fiber.StatusInternalServerError, err.Error(), nil, 0)
	}

	return utils.WriteToResponseBody(c, err, fiber.StatusOK, "Success", orderList, len(orderList))
}

func NewOrderHandler(order order.Service) Handler {
	return &handlerImpl{order: order}
}
