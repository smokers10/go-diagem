package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/infrastructure/etc"
)

type resellerController struct{}

func ResellerController() resellerController {
	return resellerController{}
}

func (rc *resellerController) IndexPage(c *fiber.Ctx) error {
	return c.Render("umum/seller/index", fiber.Map{"Logged": etc.GetLocal(c)})
}
