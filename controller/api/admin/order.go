package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type orderController struct {
	orderService domain.OrderService
}

func OrderController(order *domain.OrderService) *orderController {
	return &orderController{orderService: *order}
}

func (oc *orderController) Read(c *fiber.Ctx) error {
	response := oc.orderService.ReadAll()
	return c.Status(response.Status).JSON(response)
}

func (oc *orderController) UpdateStatus(c *fiber.Ctx) error {
	orderID := c.FormValue("order_id")
	status := c.FormValue("status")

	response := oc.orderService.UpdateStatus(orderID, status)
	return c.Status(response.Status).JSON(response)
}

func (oc *orderController) DetailOrder(c *fiber.Ctx) error {
	orderID := c.Params("order_id")

	response := oc.orderService.Detail(orderID)
	return c.Status(response.Status).JSON(response)
}
