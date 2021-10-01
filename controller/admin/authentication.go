package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type adminAuthentication struct {
	adminService domain.AdminService
}

func AdminAuthentication(adminService *domain.AdminService) adminAuthentication {
	return adminAuthentication{adminService: *adminService}
}

func (a *adminAuthentication) Login(c *fiber.Ctx) error {
	cred := domain.AdminCredentials{}

	if err := c.BodyParser(&cred); err != nil {
		panic(err)
	}

	response := a.adminService.Login(&cred)

	return c.Status(response.Status).JSON(response)
}
