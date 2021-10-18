package admin

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type sliderController struct {
	sliderService domain.SliderService
}

func SliderController(slider *domain.SliderService) sliderController {
	return sliderController{sliderService: *slider}
}

func (s *sliderController) Create(c *fiber.Ctx) error {
	req := domain.Slider{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	response := s.sliderService.Create(&req)

	return c.Status(response.Status).JSON(response)
}

func (s *sliderController) Read(c *fiber.Ctx) error {
	response := s.sliderService.Read()

	return c.Status(response.Status).JSON(response)
}

func (s *sliderController) Update(c *fiber.Ctx) error {
	req := domain.Slider{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	response := s.sliderService.Update(&req)

	return c.Status(response.Status).JSON(response)
}

func (s *sliderController) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.FormValue("id"))

	response := s.sliderService.Delete(id)

	return c.Status(response.Status).JSON(response)
}
