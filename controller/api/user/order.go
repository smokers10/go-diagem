package user

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

func (oc *orderController) Order(c *fiber.Ctx) error {
	req := domain.Order{}
	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}
	userID := c.Locals("id").(int)
	req.UserID = userID
	response := oc.orderService.Create(&req)
	return c.Status(response.Status).JSON(response)
}

func (oc *orderController) Read(c *fiber.Ctx) error {
	userID := c.Locals("id").(int)
	response := oc.orderService.ReadByUser(userID)
	return c.Status(response.Status).JSON(response)
}

func (oc *orderController) DetailOrder(c *fiber.Ctx) error {
	orderID := c.Params("order_id")

	response := oc.orderService.Detail(orderID)
	return c.Status(response.Status).JSON(response)
}
