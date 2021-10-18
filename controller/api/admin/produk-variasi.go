package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type produkVariasiController struct {
	produkVariasiService domain.ProdukVariasiService
}

func ProdukVariasiController(produkVariasi *domain.ProdukVariasiService) produkVariasiController {
	return produkVariasiController{produkVariasiService: *produkVariasi}
}

func (p *produkVariasiController) Create(c *fiber.Ctx) error {
	req := domain.ProdukVariasi{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	response := p.produkVariasiService.Create(&req)

	return c.Status(response.Status).JSON(response)
}

func (p *produkVariasiController) Update(c *fiber.Ctx) error {
	req := domain.ProdukVariasi{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	response := p.produkVariasiService.Update(&req)

	return c.Status(response.Status).JSON(response)
}

func (p *produkVariasiController) Delete(c *fiber.Ctx) error {
	produkID := c.FormValue("produk_id")
	ID := c.FormValue("id")

	response := p.produkVariasiService.Delete(produkID, ID)

	return c.Status(response.Status).JSON(response)
}
