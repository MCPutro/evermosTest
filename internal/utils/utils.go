package utils

import (
	"evermosTest/internal/handler/response"
	"github.com/gofiber/fiber/v2"
)

func WriteToResponseBody(c *fiber.Ctx, err error, statusCode int, message string, data interface{}, totalRecord int) error {
	var success = true
	if err != nil {
		success = false
	}
	return c.Status(statusCode).
		JSON(response.Response{
			Code:         statusCode,
			Success:      success,
			Message:      message,
			TotalRecords: totalRecord,
			Data:         data,
		})
}
