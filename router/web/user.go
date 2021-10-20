package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	userAPI "github.com/smokers10/go-diagem.git/controller/api/user"
	"github.com/smokers10/go-diagem.git/controller/web/user"

	"github.com/smokers10/go-diagem.git/infrastructure/middleware"
	"github.com/smokers10/go-diagem.git/infrastructure/resolver"
)

func UserWebPagge(app *fiber.App, session *session.Store, resolver *resolver.ServiceResolver) {
	// middleware init
	middlewareStrict := middleware.UserStrictWeb(session)
	middlewareVerification := middleware.UserVerificationWeb(session)
	MiddlewareGuestOnly := middleware.GuestOnly(session)

	// controller init
	authencticationController := user.AuthenticationController(resolver.UserService, session)
	homeController := user.HomeController()
	verifikasiController := user.VerificationController()

	// API Controller init
	verificationAPIController := userAPI.VerificationController(resolver.VerificationService)

	// router clustering
	parentPath := app.Group("/")

	// authentication
	parentPath.Get("/login", MiddlewareGuestOnly, authencticationController.LoginPage)
	parentPath.Post("/login", MiddlewareGuestOnly, authencticationController.Login)
	parentPath.Get("/register", MiddlewareGuestOnly, authencticationController.RegisterPage)
	parentPath.Post("/register", MiddlewareGuestOnly, authencticationController.Register)
	parentPath.Get("/logout", authencticationController.Logout)

	// verifikasi
	verifikasi := parentPath.Group("/verifikasi", middlewareVerification)
	verifikasi.Get("/", verifikasiController.VerificationPage)
	verifikasi.Post("/create", verificationAPIController.Create)
	verifikasi.Post("/verificate", verificationAPIController.Verificate)

	// home
	parentPath.Get("/home", middlewareStrict, homeController.HomePage)
}
