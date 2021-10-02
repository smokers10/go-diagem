package admin

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type mitraController struct {
	mitraService domain.MitraService
}

func MitraController(mitra *domain.MitraService) mitraController {
	return mitraController{mitraService: *mitra}
}

func (m *mitraController) Create(c *fiber.Ctx) error {
	req := domain.Mitra{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	response := m.mitraService.Create(&req)

	return c.Status(response.Status).JSON(response)
}

func (m *mitraController) Read(c *fiber.Ctx) error {
	response := m.mitraService.Read()

	return c.Status(response.Status).JSON(response)
}

func (m *mitraController) Update(c *fiber.Ctx) error {
	req := domain.Mitra{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	response := m.mitraService.Update(&req)

	return c.Status(response.Status).JSON(response)
}

func (m *mitraController) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.FormValue("id"))

	response := m.mitraService.Delete(id)

	return c.Status(response.Status).JSON(response)
}

func (m *mitraController) GetOne(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	response := m.mitraService.GetOne(id)

	return c.Status(response.Status).JSON(response)
}
