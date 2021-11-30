package admin

import "github.com/gofiber/fiber/v2"

type userAdministratifControllor struct{}

func UserAdministratifController() userAdministratifControllor {
	return userAdministratifControllor{}
}

func (uac *userAdministratifControllor) IndexPage(c *fiber.Ctx) error {
	return c.Render("admin/user-administratif/index", nil)
}

func (uac *userAdministratifControllor) TambahPage(c *fiber.Ctx) error {
	return c.Render("admin/user-administratif/form", nil)
}

func (uac *userAdministratifControllor) EditPage(c *fiber.Ctx) error {
	return c.Render("admin/user-administratif/edit", nil)
}
