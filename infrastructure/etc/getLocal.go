package etc

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/infrastructure/jwt"
)

func GetLocal(c *fiber.Ctx) jwt.Payload {
	var bind jwt.Payload

	if c.Locals("id") != nil {
		bind = jwt.Payload{
			ID:         c.Locals("id").(int),
			Email:      c.Locals("email").(string),
			Type:       c.Locals("type").(string),
			IsVerified: c.Locals("is_verified").(bool),
			IsLogged:   true,
		}
	}

	return bind
}
