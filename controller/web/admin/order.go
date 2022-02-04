package admin

import "github.com/gofiber/fiber/v2"

type orderController struct{}

func OrderController() orderController {
	return orderController{}
}

func (o *orderController) IndexPage(c *fiber.Ctx) error {
	return c.Render("admin/order/order", nil)
}

func (o *orderController) DetailPage(c *fiber.Ctx) error {
	return c.Render("admin/order/detail", nil)
}
