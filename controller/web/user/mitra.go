package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/infrastructure/etc"
)

type mitraController struct{}

func MitraController() mitraController {
	return mitraController{}
}

func (m *mitraController) IndexPage(c *fiber.Ctx) error {
	return c.Render("umum/seller/index", fiber.Map{"Logged": etc.GetLocal(c)})
}
