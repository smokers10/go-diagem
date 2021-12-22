package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type alamatOriginController struct {
	alamatOriginService domain.AlamatOriginService
}

func AlamatOriginController(alamatOrigin *domain.AlamatOriginService) *alamatOriginController {
	return &alamatOriginController{alamatOriginService: *alamatOrigin}
}

func (aoc *alamatOriginController) Read(c *fiber.Ctx) error {
	response := aoc.alamatOriginService.Read()
	return c.Status(response.Status).JSON(response)
}

func (aoc *alamatOriginController) Update(c *fiber.Ctx) error {
	req := domain.AlamatOrigin{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	response := aoc.alamatOriginService.Update(&req)
	return c.Status(response.Status).JSON(response)
}
