package admin

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type rekeningController struct {
	rekingService domain.RekeningService
}

func RekeningController(rekening *domain.RekeningService) rekeningController {
	return rekeningController{rekingService: *rekening}
}

func (r *rekeningController) Create(c *fiber.Ctx) error {
	req := domain.Rekening{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	response := r.rekingService.Create(&req)

	return c.Status(response.Status).JSON(response)
}

func (r *rekeningController) Read(c *fiber.Ctx) error {
	response := r.rekingService.Read()

	return c.Status(response.Status).JSON(response)
}

func (r *rekeningController) Update(c *fiber.Ctx) error {
	req := domain.Rekening{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	response := r.rekingService.Update(&req)

	return c.Status(response.Status).JSON(response)
}

func (r *rekeningController) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.FormValue("id"))

	response := r.rekingService.Delete(id)

	return c.Status(response.Status).JSON(response)
}
