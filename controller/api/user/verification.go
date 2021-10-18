package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type verificationController struct {
	verificationService domain.VerifikasiService
}

func VerificationController(verif *domain.VerifikasiService) verificationController {
	return verificationController{verificationService: *verif}
}

func (v *verificationController) Create(c *fiber.Ctx) error {
	req := domain.Verifikasi{}

	req.UserID = c.Locals("id").(int)

	response := v.verificationService.Create(&req)

	return c.Status(response.Status).JSON(response)
}

func (v *verificationController) Verificate(c *fiber.Ctx) error {
	req := domain.Verifikasi{}
	kode := c.FormValue("kode")

	req.UserID = c.Locals("id").(int)
	req.Kode = kode

	response := v.verificationService.Verificate(&req)

	return c.Status(response.Status).JSON(response)
}
