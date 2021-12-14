package etc

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/infrastructure/jwt"
)

// Get local untuk mendapatkan semua data claims JWT yang ada dalam c.locals
func GetLocal(c *fiber.Ctx) jwt.Payload {
	var localVals jwt.Payload

	if c.Locals("id") != nil {
		localVals = jwt.Payload{
			ID:         c.Locals("id").(int),
			Email:      c.Locals("email").(string),
			Type:       c.Locals("type").(string),
			IsVerified: c.Locals("is_verified").(bool),
			IsLogged:   true,
		}
	}

	return localVals
}
