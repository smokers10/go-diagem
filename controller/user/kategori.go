package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type kategoriController struct {
	kategoriService domain.KategoriService
}

func KategoriController(kategori *domain.KategoriService) kategoriController {
	return kategoriController{kategoriService: *kategori}
}

func (k *kategoriController) Read(c *fiber.Ctx) error {
	response := k.kategoriService.Read()
	return c.Status(response.Status).JSON(response)
}
