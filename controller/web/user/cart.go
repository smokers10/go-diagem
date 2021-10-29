package user

import "github.com/gofiber/fiber/v2"

type cartController struct{}

func CartController() cartController {
	return cartController{}
}

func (cc *cartController) IndexPage(c *fiber.Ctx) error {
	return c.Render("umum/cart/index", nil)
}

func (cc *cartController) CheckoutPage(c *fiber.Ctx) error {
	return c.Render("umum/cart/checkout", nil)
}

func (cc *cartController) PembayaranPage(c *fiber.Ctx) error {
	return c.Render("umum/cart/pembayaran", nil)
}
