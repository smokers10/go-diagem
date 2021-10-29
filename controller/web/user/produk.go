package user

import "github.com/gofiber/fiber/v2"

type produkController struct{}

func ProdukController() produkController {
	return produkController{}
}

func (p *produkController) IndexPage(c *fiber.Ctx) error {
	return c.Render("umum/produk/index", nil)
}

func (p *produkController) DetailPage(c *fiber.Ctx) error {
	return c.Render("umum/produk/detail", nil)
}

func (p *produkController) ProdukDetailPage(c *fiber.Ctx) error {
	return c.Render("umum/produk/produk_detail", nil)
}
