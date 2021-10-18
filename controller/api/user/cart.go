package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
)

type cartController struct {
	cartService domain.CartService
}

func CartController(cart *domain.CartService) cartController {
	return cartController{cartService: *cart}
}

func (cc *cartController) Read(c *fiber.Ctx) error {
	userID := c.Locals("id").(int)

	response := cc.cartService.Read(userID)

	return c.Status(response.Status).JSON(response)
}

func (cc *cartController) AddToCart(c *fiber.Ctx) error {
	req := domain.Cart{}
	userID := c.Locals("id").(int)

	if err := c.BodyParser(&req); err != nil {
		panic(err)
	}

	req.UserID = userID

	response := cc.cartService.AddtoCart(&req)

	return c.Status(response.Status).JSON(response)
}

func (cc *cartController) UpdateQuantity(c *fiber.Ctx) error {
	req := domain.CartData{}
	userID := c.Locals("id").(int)

	if err := c.BodyParser(&req); err != nil {
		panic(req)
	}

	req.UserID = userID

	response := cc.cartService.UpdateQuantity(&req)

	return c.Status(response.Status).JSON(response)
}

func (cc *cartController) DeleteCart(c *fiber.Ctx) error {
	req := domain.CartData{}
	userID := c.Locals("id").(int)

	if err := c.BodyParser(&req); err != nil {
		panic(req)
	}

	req.UserID = userID

	response := cc.cartService.Delete(&req)

	return c.Status(response.Status).JSON(response)
}
