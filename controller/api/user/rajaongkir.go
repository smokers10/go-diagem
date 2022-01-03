package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type rajaongkir struct {
	alamatService domain.AlamatService
}

func RajaOngkir(user *domain.AlamatService) rajaongkir {
	return rajaongkir{alamatService: *user}
}

func (a *rajaongkir) Provinsi(c *fiber.Ctx) error {
	response := a.alamatService.GetProvinsi()
	return c.Status(response.Status).JSON(response)
}

func (a *rajaongkir) Kota(c *fiber.Ctx) error {
	provinsiID := c.Params("provinsi_id")
	response := a.alamatService.GetKota(provinsiID)
	return c.Status(response.Status).JSON(response)
}

func (a *rajaongkir) AllKota(c *fiber.Ctx) error {
	response := a.alamatService.GetAllKota()
	return c.Status(response.Status).JSON(response)
}
