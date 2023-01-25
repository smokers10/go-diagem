package user

import "github.com/gofiber/fiber/v2"

type whoWeAreController struct{}

func WhoWeAreController() whoWeAreController {
	return whoWeAreController{}
}

func (wwa *whoWeAreController) Index(c *fiber.Ctx) error {
	return c.Render("umum/who-we-are/index", nil)
}
