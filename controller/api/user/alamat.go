package user

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type alamatController struct {
	alamatService domain.AlamatService
}

func AlamatController(user *domain.AlamatService) alamatController {
	return alamatController{alamatService: *user}
}

func (a *alamatController) Create(c *fiber.Ctx) error {
	req := domain.Alamat{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	req.UserID = c.Locals("id").(int)

	response := a.alamatService.Create(&req)

	return c.Status(response.Status).JSON(response)
}

func (a *alamatController) Read(c *fiber.Ctx) error {
	userID := c.Locals("id").(int)

	response := a.alamatService.Read(userID)

	return c.Status(response.Status).JSON(response)
}

func (a *alamatController) Update(c *fiber.Ctx) error {
	req := domain.Alamat{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	req.UserID = c.Locals("id").(int)

	response := a.alamatService.Update(&req)

	return c.Status(response.Status).JSON(response)
}

func (a *alamatController) Delete(c *fiber.Ctx) error {
	userID := c.Locals("id").(int)
	alamatID, _ := strconv.Atoi(c.FormValue("id"))

	response := a.alamatService.Delete(alamatID, userID)

	return c.Status(response.Status).JSON(response)
}

func (a *alamatController) MakeUtama(c *fiber.Ctx) error {
	userID := c.Locals("id").(int)
	alamatID, _ := strconv.Atoi(c.FormValue("id"))

	response := a.alamatService.MakeUtama(alamatID, userID)

	return c.Status(response.Status).JSON(response)
}

func (a *alamatController) Provinsi(c *fiber.Ctx) error {
	response := a.alamatService.GetProvinsi()
	return c.Status(response.Status).JSON(response)
}

func (a *alamatController) Kota(c *fiber.Ctx) error {
	provinsiID := c.Params("provinsi_id")
	response := a.alamatService.GetKota(provinsiID)
	return c.Status(response.Status).JSON(response)
}
