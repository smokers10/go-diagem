package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type produkFotoController struct {
	produkFotoService domain.ProdukFotoService
}

func ProdukFotoController(produkFoto *domain.ProdukFotoService) produkFotoController {
	return produkFotoController{produkFotoService: *produkFoto}
}

func (pc *produkFotoController) Upload(c *fiber.Ctx) error {
	req := domain.FotoProdukFile{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	response := pc.produkFotoService.Create(&req)

	return c.Status(response.Status).JSON(response)
}

func (pc *produkFotoController) MakeUtama(c *fiber.Ctx) error {
	id := c.FormValue("id")
	produkID := c.FormValue("produk_id")

	response := pc.produkFotoService.MakeUtama(id, produkID)
	return c.Status(response.Status).JSON(response)
}

func (pc *produkFotoController) Delete(c *fiber.Ctx) error {
	id := c.FormValue("id")
	path := c.FormValue("path")

	response := pc.produkFotoService.Delete(id, path)

	return c.Status(response.Status).JSON(response)
}
