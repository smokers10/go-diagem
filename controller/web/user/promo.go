package user

import "github.com/gofiber/fiber/v2"

type promoController struct{}

func PromoController() promoController {
	return promoController{}
}

func (p *promoController) IndexPage(c *fiber.Ctx) error {
	return c.Render("umum/promo/index", nil)
}

func (p *promoController) DetailPage(c *fiber.Ctx) error {
	return c.Render("umum/promo/detail", nil)
}
