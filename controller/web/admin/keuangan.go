package admin

import "github.com/gofiber/fiber/v2"

type keuanganController struct{}

func KeuanganController() keuanganController {
	return keuanganController{}
}

func (k *keuanganController) RekeningPage(c *fiber.Ctx) error {
	return c.Render("admin/keuangan/rekening", nil)
}
