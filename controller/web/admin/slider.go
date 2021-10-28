package admin

import "github.com/gofiber/fiber/v2"

type sliderController struct{}

func SliderController() sliderController {
	return sliderController{}
}

func (s *sliderController) IndexPage(c *fiber.Ctx) error {
	return c.Render("admin/slider/index", nil)
}

func (s *sliderController) TambahPage(c *fiber.Ctx) error {
	return c.Render("admin/slider/form", nil)
}

func (s *sliderController) EditPage(c *fiber.Ctx) error {
	return c.Render("admin/slider/edit", nil)
}
