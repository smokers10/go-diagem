package user

import "github.com/gofiber/fiber/v2"

type verificationController struct{}

func VerificationController() verificationController {
	return verificationController{}
}

func (vc *verificationController) VerificationPage(c *fiber.Ctx) error {
	return c.Render("umum/verifikasi", nil)
}
