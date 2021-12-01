package admin

import "github.com/gofiber/fiber/v2"

type userControllerImpl struct{}

func UserController() userControllerImpl {
	return userControllerImpl{}
}

func (uc *userControllerImpl) IndexPage(c *fiber.Ctx) error {
	return c.Render("admin/user/index", nil)
}

func (uc *userControllerImpl) DetailPage(c *fiber.Ctx) error {
	return c.Render("admin/user/profile", nil)
}
