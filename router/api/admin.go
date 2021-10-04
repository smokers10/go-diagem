package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/controller/admin"
	"github.com/smokers10/go-diagem.git/infrastructure/middleware"
	"github.com/smokers10/go-diagem.git/infrastructure/resolver"
)

func AdminAPI(app *fiber.App, resolver *resolver.ServiceResolver) {
	//middleware
	middleware := middleware.Admin()

	//controller
	authentication := admin.AdminAuthentication(resolver.AdminService)
	kategoriController := admin.KategoriController(resolver.KategoriService)
	mitraController := admin.MitraController(resolver.MitraService)
	produkController := admin.ProdukController(resolver.ProdukService)
	promoController := admin.PromoController(resolver.PromoService)

	//parent path
	parent := app.Group("/api/admin")

	//auth
	auth := parent.Group("/auth")
	auth.Post("/login", authentication.Login)

	//kategori
	kategori := parent.Group("/kategori")
	kategori.Post("/create", middleware, kategoriController.Create)
	kategori.Get("/list", middleware, kategoriController.Read)
	kategori.Put("/update", middleware, kategoriController.Update)
	kategori.Delete("/delete", middleware, kategoriController.Delete)

	//mitra
	mitra := parent.Group("/mitra")
	mitra.Post("/create", middleware, mitraController.Create)
	mitra.Get("/list", middleware, mitraController.Read)
	mitra.Get("/:id", middleware, mitraController.GetOne)
	mitra.Put("/update", middleware, mitraController.Update)
	mitra.Delete("/delete", middleware, mitraController.Delete)

	//produk
	produk := parent.Group("/produk")
	produk.Post("/create", middleware, produkController.Create)
	produk.Get("/list", middleware, produkController.Read)
	produk.Put("/update", middleware, produkController.Delete)
	produk.Post("/delete", middleware, produkController.Delete)
	produk.Get("/:id", middleware, produkController.Detail)

	//promo
	promo := parent.Group("/promo")
	promo.Post("/create", middleware, promoController.Create)
	promo.Get("/list", middleware, promoController.Read)
	promo.Get("/:id", middleware, promoController.Detail)
	promo.Put("/update", middleware, promoController.Update)
	promo.Delete("/delete", middleware, promoController.Delete)

	//
}
