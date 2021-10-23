package user

import "github.com/gofiber/fiber/v2"

type alamatController struct{}

func AlamatController() alamatController {
	return alamatController{}
}

func (ac *alamatController) AlamatPage(c *fiber.Ctx) error {
	return c.Render("umum/user/alamat", nil)
}
