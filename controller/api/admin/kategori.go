package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type kategoriController struct {
	kategoriService domain.KategoriService
}

func KategoriController(kategori *domain.KategoriService) *kategoriController {
	return &kategoriController{kategoriService: *kategori}
}

func (k *kategoriController) Create(c *fiber.Ctx) error {
	req := domain.Kategori{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	response := k.kategoriService.Create(&req)

	return c.Status(response.Status).JSON(response)
}

func (k *kategoriController) Read(c *fiber.Ctx) error {
	response := k.kategoriService.Read()

	return c.Status(response.Status).JSON(response)
}

func (k *kategoriController) Update(c *fiber.Ctx) error {
	req := domain.Kategori{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	response := k.kategoriService.Update(&req)

	return c.Status(response.Status).JSON(response)
}

func (k *kategoriController) UpdateCover(c *fiber.Ctx) error {
	req := domain.Kategori{}

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	response := k.kategoriService.UpdateCover(&req)

	return c.Status(response.Status).JSON(response)
}

func (k *kategoriController) Delete(c *fiber.Ctx) error {
	id := c.FormValue("id")

	response := k.kategoriService.Delete(id)

	return c.Status(response.Status).JSON(response)
}

func (k *kategoriController) Detail(c *fiber.Ctx) error {
	id := c.Params("id")

	response := k.kategoriService.Detail(id)

	return c.Status(response.Status).JSON(response)
}
