package user

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type alamatController struct {
	alamarService domain.AlamatService
}

func AlamatController(user *domain.AlamatService) alamatController {
	return alamatController{alamarService: *user}
}

func (a *alamatController) Create(c *fiber.Ctx) error {
	req := domain.Alamat{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	req.UserID = c.Locals("id").(int)

	response := a.alamarService.Create(&req)

	return c.Status(response.Status).JSON(response)
}

func (a *alamatController) Read(c *fiber.Ctx) error {
	userID := c.Locals("id").(int)

	response := a.alamarService.Read(userID)

	return c.Status(response.Status).JSON(response)
}

func (a *alamatController) Update(c *fiber.Ctx) error {
	req := domain.Alamat{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	req.UserID = c.Locals("id").(int)
	req.KDPos = c.FormValue("kd_pos")

	response := a.alamarService.Update(&req)

	return c.Status(response.Status).JSON(response)
}

func (a *alamatController) Delete(c *fiber.Ctx) error {
	userID := c.Locals("id").(int)
	alamatID, _ := strconv.Atoi(c.FormValue("id"))

	response := a.alamarService.Delete(alamatID, userID)

	return c.Status(response.Status).JSON(response)
}

func (a *alamatController) MakeUtama(c *fiber.Ctx) error {
	userID := c.Locals("id").(int)
	alamatID, _ := strconv.Atoi(c.FormValue("id"))

	response := a.alamarService.MakeUtama(alamatID, userID)

	return c.Status(response.Status).JSON(response)
}
