package user

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type checkoutControllerImpl struct {
	checkoutService domain.CheckoutService
}

func CheckoutController(checkout *domain.CheckoutService) checkoutControllerImpl {
	return checkoutControllerImpl{checkoutService: *checkout}
}

func (cc *checkoutControllerImpl) Checkout(c *fiber.Ctx) error {
	userID := c.Locals("id").(int)
	alamatID, _ := strconv.Atoi(c.FormValue("alamat_id"))
	req := domain.Checkout{
		UserID:   userID,
		AlamatID: alamatID,
	}

	response := cc.checkoutService.Checkout(&req)

	return c.Status(response.Status).JSON(response)
}
