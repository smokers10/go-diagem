package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type resetPasswordController struct {
	resetPasswordService domain.PasswordResetsService
}

func ResetPasswordController(resetPW *domain.PasswordResetsService) *resetPasswordController {
	return &resetPasswordController{resetPasswordService: *resetPW}
}

func (rpc *resetPasswordController) ForgotPasswordPage(c *fiber.Ctx) error {
	return c.Render("umum/forgot-password/send", nil)
}

func (rpc *resetPasswordController) ResetPasswordPage(c *fiber.Ctx) error {
	return c.Render("umum/forgot-password/reset", nil)
}

func (rpc *resetPasswordController) Create(c *fiber.Ctx) error {
	email := c.FormValue("email")
	response := rpc.resetPasswordService.Create(email)
	return c.Status(response.Status).JSON(response)
}

func (rpc *resetPasswordController) Reset(c *fiber.Ctx) error {
	req := domain.PasswordResets{}
	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}
	response := rpc.resetPasswordService.Reset(&req)
	return c.Status(response.Status).JSON(response)
}
