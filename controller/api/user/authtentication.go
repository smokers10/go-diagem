package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type userAuthController struct {
	userService domain.UserService
}

func UserAuthController(user *domain.UserService) userAuthController {
	return userAuthController{userService: *user}
}

func (u *userAuthController) Login(c *fiber.Ctx) error {
	credentialReq := domain.UserCredential{}

	if err := c.BodyParser(&credentialReq); err != nil {
		panic(err)
	}

	response := u.userService.Login(&credentialReq)

	return c.Status(response.Status).JSON(response)
}

func (u *userAuthController) Register(c *fiber.Ctx) error {
	req := domain.UserBasicData{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	response := u.userService.Registrasi(&req)

	return c.Status(response.Status).JSON(response)
}

func (u *userAuthController) LoginPage(c *fiber.Ctx) error {
	return c.Render("", nil)
}
