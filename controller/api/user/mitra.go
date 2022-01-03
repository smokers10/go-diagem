package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type mitraController struct {
	mitraService domain.MitraService
}

func MitraController(mitra *domain.MitraService) *mitraController {
	return &mitraController{mitraService: *mitra}
}

func (mc *mitraController) Get(c *fiber.Ctx) error {
	req := domain.Mitra{}
	req.Nama = c.Query("nama")
	req.SellerID = c.Query("nama")
	req.KotaID = c.Query("kota")

	res := mc.mitraService.ReadWithFilter(&req)

	return c.Status(res.Status).JSON(res)
}
