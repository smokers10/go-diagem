package admin

import "github.com/gofiber/fiber/v2"

type mitraController struct{}

func MitraController() mitraController {
	return mitraController{}
}

func (m *mitraController) IndexPage(c *fiber.Ctx) error {
	return c.Render("admin/mitra/index", nil)
}

func (m *mitraController) TambahPage(c *fiber.Ctx) error {
	return c.Render("admin/mitra/tambah", nil)
}

func (m *mitraController) EditPage(c *fiber.Ctx) error {
	return c.Render("admin/mitra/edit", nil)
}
