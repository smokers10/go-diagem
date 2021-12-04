package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/smokers10/go-diagem.git/infrastructure/jwt"
)

func UserStrictWeb(session *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := session.Get(c)
		if err != nil {
			panic(err)
		}

		// ambil token dari session
		token := sess.Get("token")

		// check token
		if token == nil {
			return c.Redirect("/login")
		}

		// verifikasi token
		claims := jwt.Verify(token.(string))

		// check authorisasi
		if claims.Type != "user" {
			return c.Redirect("/")
		}

		// check verifikasi email
		if !claims.IsVerified {
			return c.Redirect("/verifikasi")
		}

		// assign claims ke locals
		c.Locals("id", claims.ID)
		c.Locals("email", claims.Email)
		c.Locals("type", claims.Type)
		c.Locals("is_verified", claims.IsVerified)

		return c.Next()
	}
}

func UserNotSoStrictWeb(session *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := session.Get(c)
		if err != nil {
			panic(err)
		}

		// ambil token dari session
		token := sess.Get("token")

		if token != nil {
			// verifikasi token
			claims := jwt.Verify(token.(string))

			// check authorisasi
			if claims.Type != "user" {
				return c.Redirect("/")
			}

			// check verifikasi email
			if !claims.IsVerified {
				return c.Redirect("/verifikasi")
			}

			// assign claims ke locals
			c.Locals("id", claims.ID)
			c.Locals("email", claims.Email)
			c.Locals("type", claims.Type)
			c.Locals("is_verified", claims.IsVerified)

			return c.Next()
		}

		return c.Next()
	}
}

func UserVerificationWeb(session *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := session.Get(c)
		if err != nil {
			panic(err)
		}

		// ambil token dari session
		token := sess.Get("token")

		// check token
		if token == nil {
			return c.Redirect("/login")
		}

		// verifikasi token
		claims := jwt.Verify(token.(string))

		// check authorisasi
		if claims.Type != "user" {
			return c.Redirect("/")
		}

		// check verifikasi email, redirect ke home jika sudah terverifikasi
		if claims.IsVerified {
			return c.Redirect("/home")
		}

		// assign claims ke locals
		c.Locals("id", claims.ID)
		c.Locals("email", claims.Email)
		c.Locals("type", claims.Type)
		c.Locals("is_verified", claims.IsVerified)

		return c.Next()
	}
}

func GuestOnly(session *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := session.Get(c)
		if err != nil {
			panic(err)
		}

		// ambil token dari session
		token := sess.Get("token")

		// check token
		if token == nil {
			return c.Next()
		}

		// verifikasi token
		claims := jwt.Verify(token.(string))

		// jika type logged user = user/umum redirect ke home
		if claims.Type == "user" {
			return c.Redirect("/home")
		}

		if claims.Type == "admin" || claims.Type == "super admin" || claims.Type == "marketing" {
			return c.Redirect("/admin/home")
		}

		return c.Next()
	}
}

func AdminWeb(session *session.Store, adminType ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := session.Get(c)
		if err != nil {
			panic(err)
		}
		var isAllowed bool

		// get token
		token := sess.Get("token")

		// check token
		if token == nil {
			return c.Redirect("/admin/login")
		}

		// verifikasi token
		claims := jwt.Verify(token.(string))

		// check authorisasi
		for _, allowed := range adminType {
			if claims.Type == allowed {
				isAllowedMemAddress := &isAllowed
				*isAllowedMemAddress = true
				break
			}
		}

		if !isAllowed {
			return c.Redirect("/admin/home")
		}

		// assign claims ke locals
		c.Locals("id", claims.ID)
		c.Locals("email", claims.Email)
		c.Locals("type", claims.Type)
		c.Locals("is_verified", claims.IsVerified)

		return c.Next()
	}
}
