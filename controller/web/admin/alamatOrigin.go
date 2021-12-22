package admin

import "github.com/gofiber/fiber/v2"

type alamatOriginController struct{}

func AlamatOriginController() *alamatOriginController {
	return &alamatOriginController{}
}

func (aoc *alamatOriginController) IndexPage(c *fiber.Ctx) error {
	return c.Render("admin/alamat", nil)
}
