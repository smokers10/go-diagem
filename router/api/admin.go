package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/controller/admin"
	"github.com/smokers10/go-diagem.git/infrastructure/resolver"
)

func AdminAPI(app *fiber.App, resolver *resolver.ServiceResolver) {
	//controller
	authentication := admin.AdminAuthentication(resolver.AdminService)

	//parent path
	parent := app.Group("/api/admin")

	parent.Post("/auth/login", authentication.Login)

}
