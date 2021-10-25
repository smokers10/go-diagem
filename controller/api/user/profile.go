package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type profileControllerImpl struct {
	userService domain.UserService
}

func ProfileController(user *domain.UserService) profileControllerImpl {
	return profileControllerImpl{userService: *user}
}

func (p *profileControllerImpl) GetProfile(c *fiber.Ctx) error {
	id := c.Locals("id").(int)

	response := p.userService.GetProfile(id)

	return c.Status(response.Status).JSON(response)
}

func (p *profileControllerImpl) Update(c *fiber.Ctx) error {
	req := domain.UserBasicData{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	req.ID = c.Locals("id").(int)

	response := p.userService.Update(&req)

	return c.Status(response.Status).JSON(response)
}
