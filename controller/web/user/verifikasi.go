package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/smokers10/go-diagem.git/domain"
)

type verificationController struct {
	verifikasiService domain.VerifikasiService
	session           *session.Store
}

func VerificationController(verification *domain.VerifikasiService, session *session.Store) verificationController {
	return verificationController{
		verifikasiService: *verification,
		session:           session,
	}
}

func (vc *verificationController) VerificationPage(c *fiber.Ctx) error {
	return c.Render("umum/verifikasi", nil)
}

func (vc *verificationController) Verificate(c *fiber.Ctx) error {
	req := domain.Verifikasi{}
	kode := c.FormValue("kode")
	sess, err := vc.session.Get(c)
	if err != nil {
		panic(err)
	}

	req.UserID = c.Locals("id").(int)
	req.Kode = kode

	response := vc.verifikasiService.Verificate(&req)
	if response.Success {
		sess.Set("token", fmt.Sprintf("Bearer %s", response.Token))
		sess.Save()
	}

	return c.Status(response.Status).JSON(response)
}
