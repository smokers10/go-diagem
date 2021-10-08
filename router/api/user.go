package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/controller/user"
	"github.com/smokers10/go-diagem.git/infrastructure/middleware"
	"github.com/smokers10/go-diagem.git/infrastructure/resolver"
)

func UserAPI(app *fiber.App, resolver *resolver.ServiceResolver) {
	//middleware
	middlewareStrict := middleware.UserStrict()
	middlewareNonStrict := middleware.UserNonStrict()

	// controller
	authController := user.UserAuthController(resolver.UserService)
	verificationController := user.VerificationController(resolver.VerificationService)

	// parent path
	parentPath := app.Group("/api/user")

	// authentication
	auth := parentPath.Group("/auth")
	auth.Post("/register", authController.Register)
	auth.Post("/login", authController.Login)

	// verification
	verification := parentPath.Group("/verification")
	verification.Post("/create", middlewareNonStrict, verificationController.Create)
	verification.Post("/verificate", middlewareNonStrict, verificationController.Verificate)

	parentPath.Get("/", middlewareStrict, func(c *fiber.Ctx) error {
		return c.JSON("Hey there wellcome")
	})
}
