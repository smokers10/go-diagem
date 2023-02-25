package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/controller/api/user"
	"github.com/smokers10/go-diagem.git/infrastructure/middleware"
	"github.com/smokers10/go-diagem.git/infrastructure/resolver"
)

func UserAPI(app *fiber.App, resolver *resolver.ServiceResolver) {
	// middleware
	middlewareStrict := middleware.UserStrict()
	middlewareNonStrict := middleware.UserNonStrict()

	// controller
	authController := user.UserAuthController(resolver.UserService)
	cartController := user.CartController(resolver.CartService)
	verificationController := user.VerificationController(resolver.VerificationService)
	alamatController := user.AlamatController(resolver.AlamatService)
	kategoriController := user.KategoriController(resolver.KategoriService)
	produkController := user.ProdukController(resolver.ProdukService)
	checkoutController := user.CheckoutController(resolver.CheckoutService)
	profileController := user.ProfileController(resolver.UserService)

	// parent path
	parentPath := app.Group("/api/user")

	// authentication
	auth := parentPath.Group("/auth")
	auth.Post("/register", authController.Register)
	auth.Post("/login", authController.Login)

	// verification
	verification := parentPath.Group("/verification", middlewareNonStrict)
	verification.Post("/create", verificationController.Create)
	verification.Post("/verificate", verificationController.Verificate)

	// alamat
	alamat := parentPath.Group("/alamat", middlewareStrict)
	alamat.Post("/create", alamatController.Create)
	alamat.Get("/read", alamatController.Read)
	alamat.Put("/update", alamatController.Update)
	alamat.Delete("/delete", alamatController.Delete)
	alamat.Put("/make-utama", alamatController.MakeUtama)

	// kategori
	kategori := parentPath.Group("/kategori")
	kategori.Get("/read", kategoriController.Read)

	// produk
	produk := parentPath.Group("/produk")
	produk.Get("/", produkController.Read)
	produk.Get("/:slug", produkController.Detail)

	// cart
	cart := parentPath.Group("/cart", middlewareStrict)
	cart.Get("/read", cartController.Read)
	cart.Post("/add-to-cart", cartController.AddToCart)
	cart.Put("/update-quantity", cartController.UpdateQuantity)
	cart.Delete("/delete", cartController.DeleteCart)

	// profile
	profile := parentPath.Group("/profile", middlewareStrict)
	profile.Get("/", profileController.GetProfile)
	profile.Put("/update", profileController.Update)

	// checkout
	checkout := parentPath.Group("/checkout", middlewareStrict)
	checkout.Post("/", checkoutController.Checkout)

	parentPath.Get("/", middlewareStrict, func(c *fiber.Ctx) error {
		return c.JSON("Hey there wellcome")
	})

	// Slider
	sliderController := user.SliderController(resolver.SliderService) 
	slider := parentPath.Group("/slider")
	slider.Get("/", sliderController.Read)
}
