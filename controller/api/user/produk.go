package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type produkControllerImpl struct {
	produkService domain.ProdukService
}

func ProdukController(produk *domain.ProdukService) produkControllerImpl {
	return produkControllerImpl{produkService: *produk}
}

func (p *produkControllerImpl) Read(c *fiber.Ctx) error {
	filter := domain.ProdukFilter{}
	if err := c.BodyParser(&filter); err != nil {
		panic(filter)
	}

	response := p.produkService.Read(&filter)
	return c.Status(response.Status).JSON(response)
}

func (p *produkControllerImpl) Detail(c *fiber.Ctx) error {
	slug := c.Params("slug")
	response := p.produkService.DetailBySlug(slug)
	return c.Status(response.Status).JSON(response)
}

func (p *produkControllerImpl) GetPopular(c *fiber.Ctx) error {
	response := p.produkService.GetPopular()
	return c.Status(response.Status).JSON(response)
}
