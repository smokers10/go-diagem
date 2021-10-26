package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/smokers10/go-diagem.git/controller/web/admin"
	"github.com/smokers10/go-diagem.git/infrastructure/middleware"
	"github.com/smokers10/go-diagem.git/infrastructure/resolver"
)

func AdminWebPage(app *fiber.App, session *session.Store, resolver *resolver.ServiceResolver) {
	// middleware
	guestOnly := middleware.GuestOnly(session)

	// controller init
	authenticationController := admin.AdminAuthenticationController(resolver.AdminService, session)
	homeController := admin.HomeController()

	// controller API ini

	// router clustering
	adminParentPath := app.Group("/admin")

	// authenctication
	adminParentPath.Get("/login", guestOnly, authenticationController.LoginPage)
	adminParentPath.Post("/login", guestOnly, authenticationController.Login)
	adminParentPath.Get("/logout", authenticationController.Logout)

	// home
	adminParentPath.Get("/home", middleware.AdminWeb(session, "admin", "super admin", "marketing"), homeController.HomePage)
}
