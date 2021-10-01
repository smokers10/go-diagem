package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
	"github.com/smokers10/go-diagem.git/infrastructure/jwt"
)

func Admin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := string(c.Request().Header.Peek("Authorization"))
		res := domain.Response{}

		// check token
		if token == "" {
			res.Message = "akses token harus di sediakan"
			res.Status = http.StatusUnauthorized
			return c.Status(http.StatusUnauthorized).JSON(res)
		}

		// verifikasi token
		claims := jwt.Verify(token)

		// check authorisasi
		if claims.Type != "admin" {
			res.Message = "hanya untuk akses ter-authorisasi"
			res.Status = http.StatusUnauthorized
			return c.Status(http.StatusUnauthorized).JSON(res)
		}

		// assign claims ke locals
		c.Locals("id", claims.ID)
		c.Locals("email", claims.ID)
		c.Locals("type", claims.Type)

		return c.Next()
	}
}
