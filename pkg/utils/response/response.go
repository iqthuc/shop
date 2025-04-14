package response

import (
	errs "shop/pkg/utils/errors"
	"shop/pkg/utils/messages"

	"github.com/gofiber/fiber/v2"
)

type APIResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Status  int    `json:"status"`
}

func JsonResponse(c *fiber.Ctx, response APIResponse) error {
	return c.Status(response.Code).JSON(response)
}

func ErrorJson(c *fiber.Ctx, err errs.AppError, code int) error {
	response := APIResponse{
		Message: err.Error(),
		Code:    code,
		Data:    nil,
		Status:  0,
	}
	return JsonResponse(c, response)
}

func SuccessJson(c *fiber.Ctx, data any, message messages.AppMessage) error {
	response := APIResponse{
		Message: string(message),
		Code:    fiber.StatusOK,
		Data:    data,
		Status:  1,
	}
	return JsonResponse(c, response)
}
