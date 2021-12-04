package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/infrastructure/etc"
)

type produkController struct{}

func ProdukController() produkController {
	return produkController{}
}

func (p *produkController) IndexPage(c *fiber.Ctx) error {
	return c.Render("umum/produk/index", fiber.Map{"Logged": etc.GetLocal(c)})
}

func (p *produkController) DetailPage(c *fiber.Ctx) error {
	return c.Render("umum/produk/detail", fiber.Map{"Logged": etc.GetLocal(c)})
}

func (p *produkController) ProdukDetailPage(c *fiber.Ctx) error {
	return c.Render("umum/produk/produk_detail", fiber.Map{"Logged": etc.GetLocal(c)})
}
