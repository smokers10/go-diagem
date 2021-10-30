package user

import "github.com/gofiber/fiber/v2"

type kategoriController struct{}

func KategoriController() kategoriController {
	return kategoriController{}
}

func (k *kategoriController) IndexPage(c *fiber.Ctx) error {
	return c.Render("umum/kategori/index", nil)
}

func (k *kategoriController) DetailPage(c *fiber.Ctx) error {
	return c.Render("umum/kategori/kategori_detail", nil)
}
