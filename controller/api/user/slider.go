package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type sliderController struct {
	sliderService domain.SliderService
}

func SliderController(sc *domain.SliderService) sliderController {
	return sliderController{sliderService: *sc}
}

func (sc *sliderController) Read(c *fiber.Ctx) error {
	response := sc.sliderService.ReadPublished()

	return c.Status(response.Status).JSON(response)
}