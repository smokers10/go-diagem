package admin

import "github.com/gofiber/fiber/v2"

type adminAuthenticationController struct{}

func AdminAuthenticationController() adminAuthenticationController {
	return adminAuthenticationController{}
}

func (aac *adminAuthenticationController) LoginPage(c *fiber.Ctx) error {
	return c.Render("admin/auth/login", nil)
}
