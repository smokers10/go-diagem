package user

import "github.com/gofiber/fiber/v2"

type orderController struct{}

func OrderController() *orderController {
	return &orderController{}
}

func (oc *orderController) IndexPage(c *fiber.Ctx) error {
	return c.Render("umum/order/pesanan", nil)
}

func (oc *orderController) Invoice(c *fiber.Ctx) error {
	return c.Render("umum/order/invoice", nil)
}
