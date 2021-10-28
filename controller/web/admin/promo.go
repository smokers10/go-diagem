package admin

import "github.com/gofiber/fiber/v2"

type promoController struct{}

func PromoController() promoController {
	return promoController{}
}

func (p *promoController) IndexPage(c *fiber.Ctx) error {
	return c.Render("admin/promo/index", nil)
}

func (p *promoController) TambahPage(c *fiber.Ctx) error {
	return c.Render("admin/promo/form", nil)
}

func (p *promoController) EditPage(c *fiber.Ctx) error {
	return c.Render("admin/promo/edit", nil)
}
