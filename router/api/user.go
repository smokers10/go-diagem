package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/controller/user"
	"github.com/smokers10/go-diagem.git/infrastructure/middleware"
	"github.com/smokers10/go-diagem.git/infrastructure/resolver"
)

func UserAPI(app *fiber.App, resolver *resolver.ServiceResolver) {
	// middleware
	middlewareStrict := middleware.UserStrict()
	middlewareNonStrict := middleware.UserNonStrict()

	// controller
	authController := user.UserAuthController(resolver.UserService)
	verificationController := user.VerificationController(resolver.VerificationService)
	alamatController := user.AlamatController(resolver.AlamatService)
	kategoriController := user.KategoriController(resolver.KategoriService)
	produkController := user.ProdukController(resolver.ProdukService)

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

	// alamat
	alamat := parentPath.Group("/alamat")
	alamat.Post("/create", middlewareStrict, alamatController.Create)
	alamat.Get("/read", middlewareStrict, alamatController.Read)
	alamat.Put("/update", middlewareStrict, alamatController.Update)
	alamat.Delete("/delete", middlewareStrict, alamatController.Delete)

	// kategori
	kategori := parentPath.Group("/kategori")
	kategori.Get("/read", kategoriController.Read)

	// produk
	produk := parentPath.Group("/produk")
	produk.Get("/", produkController.Read)
	produk.Get("/:slug", produkController.Detail)

	parentPath.Get("/", middlewareStrict, func(c *fiber.Ctx) error {
		return c.JSON("Hey there wellcome")
	})
}
