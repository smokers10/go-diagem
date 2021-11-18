package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type produkController struct {
	produkService domain.ProdukService
}

func ProdukController(produkServ *domain.ProdukService) produkController {
	return produkController{produkService: *produkServ}
}

func (p *produkController) Create(c *fiber.Ctx) error {
	req := domain.Produk{}
	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	// call service produk
	p.produkService.Create(&req)

	return c.Status(200).JSON(req)
}

func (p *produkController) Read(c *fiber.Ctx) error {
	filter := domain.ProdukFilter{}
	if err := c.QueryParser(&filter); err != nil {
		panic(err)
	}

	response := p.produkService.Read(&filter)

	return c.Status(response.Status).JSON(response)
}

func (p *produkController) Update(c *fiber.Ctx) error {
	req := domain.Produk{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	response := p.produkService.Update(&req)

	return c.Status(response.Status).JSON(response)
}

func (p *produkController) Delete(c *fiber.Ctx) error {
	id := c.FormValue("id")

	response := p.produkService.Delete(id)

	return c.Status(response.Status).JSON(response)
}

func (p *produkController) Detail(c *fiber.Ctx) error {
	id := c.Params("id")

	response := p.produkService.Detail(id)

	return c.Status(response.Status).JSON(response)
}
