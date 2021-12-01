package admin

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type userControllerImpl struct {
	userService domain.UserService
}

func UserController(user *domain.UserService) userControllerImpl {
	return userControllerImpl{userService: *user}
}

func (uc *userControllerImpl) ReadAll(c *fiber.Ctx) error {
	response := uc.userService.Read()
	return c.Status(response.Status).JSON(response)
}

func (uc *userControllerImpl) Detail(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	response := uc.userService.Detail(id)
	return c.Status(response.Status).JSON(response)
}
