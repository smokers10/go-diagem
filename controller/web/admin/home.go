package admin

import "github.com/gofiber/fiber/v2"

type homeController struct{}

func HomeController() homeController {
	return homeController{}
}

func (h *homeController) HomePage(c *fiber.Ctx) error {
	return c.Render("admin/home", nil)
}
