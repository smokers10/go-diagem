package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	userAPI "github.com/smokers10/go-diagem.git/controller/api/user"
	"github.com/smokers10/go-diagem.git/controller/web/user"

	"github.com/smokers10/go-diagem.git/infrastructure/middleware"
	"github.com/smokers10/go-diagem.git/infrastructure/resolver"
)

func UserWebPage(app *fiber.App, session *session.Store, resolver *resolver.ServiceResolver) {
	// middleware init
	middlewareStrict := middleware.UserStrictWeb(session)
	middlewareVerification := middleware.UserVerificationWeb(session)
	middlewareGuestOnly := middleware.GuestOnly(session)

	// controller init
	authencticationController := user.AuthenticationController(resolver.UserService, session)
	homeController := user.HomeController()
	verifikasiController := user.VerificationController(resolver.VerificationService, session)
	alamatController := user.AlamatController()

	// API Controller init
	verificationAPIController := userAPI.VerificationController(resolver.VerificationService)
	alamatAPIController := userAPI.AlamatController(resolver.AlamatService)

	// router clustering
	parentPath := app.Group("/")

	// authentication
	parentPath.Get("/login", middlewareGuestOnly, authencticationController.LoginPage)
	parentPath.Post("/login", middlewareGuestOnly, authencticationController.Login)
	parentPath.Get("/register", middlewareGuestOnly, authencticationController.RegisterPage)
	parentPath.Post("/register", middlewareGuestOnly, authencticationController.Register)
	parentPath.Get("/logout", authencticationController.Logout)

	// verifikasi
	verifikasi := parentPath.Group("/verifikasi", middlewareVerification)
	verifikasi.Get("/", verifikasiController.VerificationPage)
	verifikasi.Post("/create", verificationAPIController.Create)
	verifikasi.Post("/verificate", verifikasiController.Verificate)

	// home
	parentPath.Get("/home", middlewareStrict, homeController.HomePage)

	// alamat
	alamat := parentPath.Group("/alamat", middlewareStrict)
	alamat.Get("/", alamatController.AlamatPage)
	alamat.Post("/create", alamatAPIController.Create)
	alamat.Get("/read", alamatAPIController.Read)
	alamat.Put("/update", alamatAPIController.Update)
	alamat.Delete("/Delete", alamatAPIController.Delete)

	// index page
	parentPath.Get("/", middlewareGuestOnly, func(c *fiber.Ctx) error {
		return c.Render("home", nil)
	})
}
