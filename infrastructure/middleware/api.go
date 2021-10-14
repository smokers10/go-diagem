package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/domain"
	"github.com/smokers10/go-diagem.git/infrastructure/jwt"
)

func Admin(accessType string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := string(c.Request().Header.Peek("Authorization"))
		res := domain.Response{}

		// check token
		if token == "" {
			res.Message = "akses token harus disertakan"
			res.Status = http.StatusUnauthorized
			return c.Status(http.StatusUnauthorized).JSON(res)
		}

		// verifikasi token
		claims := jwt.Verify(token)

		// check authorisasi
		if claims.Type != accessType {
			res.Message = "hak authorisasi administratif tidak valid"
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

func UserStrict() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := string(c.Request().Header.Peek("Authorization"))
		res := domain.Response{}

		// check token
		if token == "" {
			res.Message = "akses token harus disertakan"
			res.Status = http.StatusUnauthorized
			return c.Status(http.StatusUnauthorized).JSON(res)
		}

		// verifikasi token
		claims := jwt.Verify(token)

		// check authorisasi
		if claims.Type != "user" {
			res.Message = "hak authorisasi tidak sesuai"
			res.Status = http.StatusUnauthorized
			return c.Status(res.Status).JSON(res)
		}

		// check verifikasi email
		if !claims.IsVerified {
			res.Message = "email akun belum terverifikasi"
			res.Status = http.StatusUnauthorized
			return c.Status(res.Status).JSON(res)
		}

		// assign claims ke locals
		c.Locals("id", claims.ID)
		c.Locals("email", claims.ID)
		c.Locals("type", claims.Type)

		return c.Next()
	}
}

func UserNonStrict() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := string(c.Request().Header.Peek("Authorization"))
		res := domain.Response{}

		// check token
		if token == "" {
			res.Message = "akses token harus disertakan"
			res.Status = http.StatusUnauthorized
			return c.Status(http.StatusUnauthorized).JSON(res)
		}

		// verifikasi token
		claims := jwt.Verify(token)

		// check authorisasi
		if claims.Type != "user" {
			res.Message = "hak authorisasi tidak sesuai"
			res.Status = http.StatusUnauthorized
			return c.Status(res.Status).JSON(res)
		}

		// assign claims ke locals
		c.Locals("id", claims.ID)
		c.Locals("email", claims.ID)
		c.Locals("type", claims.Type)

		return c.Next()
	}
}
