package user

import "github.com/gofiber/fiber/v2"

type resellerController struct{}

func ResellerController() resellerController {
	return resellerController{}
}

func (rc *resellerController) IndexPage(c *fiber.Ctx) error {
	return c.Render("umum/seller/index", nil)
}
