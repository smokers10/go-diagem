package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type bankController struct {
	bankServce domain.BankService
}

func BankController(bank *domain.BankService) *bankController {
	return &bankController{bankServce: *bank}
}

func (bc *bankController) Update(c *fiber.Ctx) error {
	req := domain.Bank{}
	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	response := bc.bankServce.Update(&req)
	return c.Status(response.Status).JSON(response)
}

func (bc *bankController) Read(c *fiber.Ctx) error {
	response := bc.bankServce.Read()
	return c.Status(response.Status).JSON(response)
}
