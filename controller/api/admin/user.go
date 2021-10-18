package admin

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type userAdminController struct {
	adminService domain.AdminService
}

func UserAdminController(admin *domain.AdminService) userAdminController {
	return userAdminController{adminService: *admin}
}

func (uac *userAdminController) Create(c *fiber.Ctx) error {
	req := domain.Admin{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	response := uac.adminService.Create(&req)

	return c.Status(response.Status).JSON(response)
}

func (uac *userAdminController) Read(c *fiber.Ctx) error {
	id := c.Locals("id").(int)

	response := uac.adminService.Read(id)

	return c.Status(response.Status).JSON(response)
}

func (uac *userAdminController) Update(c *fiber.Ctx) error {
	req := domain.Admin{}
	id := c.Locals("id").(int)
	roles := c.Locals("type").(string)

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	response := uac.adminService.Update(&req, id, roles)

	return c.Status(response.Status).JSON(response)
}

func (uac *userAdminController) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.FormValue("id"))

	response := uac.adminService.Delete(id)

	return c.Status(response.Status).JSON(response)
}
