package user

import "github.com/gofiber/fiber/v2"

type mitraController struct{}

func MitraController() mitraController {
	return mitraController{}
}

func (m *mitraController) IndexPage(c *fiber.Ctx) error {
	return c.Render("umum/seller/index", nil)
}
