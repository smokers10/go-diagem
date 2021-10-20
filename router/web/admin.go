package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/controller/web/admin"
)

func AdminWebPage(app *fiber.App) {
	// controller init
	authenticationController := admin.AdminAuthenticationController()

	// router clustering
	parentPath := app.Group("/admin")

	// authenctication
	parentPath.Get("/login", authenticationController.LoginPage)
}
