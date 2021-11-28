package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type promoController struct {
	promoService domain.PromoService
}

func PromoController(promo *domain.PromoService) promoController {
	return promoController{promoService: *promo}
}

func (p *promoController) Create(c *fiber.Ctx) error {
	req := domain.Promo{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	response := p.promoService.Create(&req)

	return c.Status(response.Status).JSON(response)
}

func (p *promoController) Read(c *fiber.Ctx) error {
	response := p.promoService.Read()

	return c.Status(response.Status).JSON(response)
}

func (p *promoController) Update(c *fiber.Ctx) error {
	req := domain.Promo{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	response := p.promoService.Update(&req)

	return c.Status(response.Status).JSON(response)
}

func (p *promoController) Delete(c *fiber.Ctx) error {
	id := c.FormValue("id")

	response := p.promoService.Delete(id)

	return c.Status(response.Status).JSON(response)
}

func (p *promoController) Detail(c *fiber.Ctx) error {
	id := c.Params("id")

	response := p.promoService.Detail(id)

	return c.Status(response.Status).JSON(response)
}
