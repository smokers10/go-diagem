package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/controller/api/admin"
	"github.com/smokers10/go-diagem.git/infrastructure/middleware"
	"github.com/smokers10/go-diagem.git/infrastructure/resolver"
)

func AdminAPI(app *fiber.App, resolver *resolver.ServiceResolver) {
	// controller
	authentication := admin.AdminAuthentication(resolver.AdminService)
	kategoriController := admin.KategoriController(resolver.KategoriService)
	mitraController := admin.MitraController(resolver.MitraService)
	produkController := admin.ProdukController(resolver.ProdukService)
	produkVariasiController := admin.ProdukVariasiController(resolver.ProdukVariasiService)
	promoController := admin.PromoController(resolver.PromoService)
	rekeningController := admin.RekeningController(resolver.RekeningService)
	sliderController := admin.SliderController(resolver.SliderService)
	userAdminController := admin.UserAdminController(resolver.AdminService)

	// parent path
	parent := app.Group("/api/admin")

	// auth
	auth := parent.Group("/auth")
	auth.Post("/login", authentication.Login)

	// kategori
	kategori := parent.Group("/kategori", middleware.Admin("marketing", "super admin"))
	kategori.Post("/create", kategoriController.Create)
	kategori.Get("/list", kategoriController.Read)
	kategori.Put("/update", kategoriController.Update)
	kategori.Delete("/delete", kategoriController.Delete)

	// mitra
	mitra := parent.Group("/mitra", middleware.Admin("admin", "super admin"))
	mitra.Post("/create", mitraController.Create)
	mitra.Get("/list", mitraController.Read)
	mitra.Get("/:id", mitraController.GetOne)
	mitra.Put("/update", mitraController.Update)
	mitra.Delete("/delete", mitraController.Delete)

	// produk
	produk := parent.Group("/produk", middleware.Admin("marketing", "super admin"))
	produk.Post("/create", produkController.Create)
	produk.Get("/list", produkController.Read)
	produk.Put("/update", produkController.Update)
	produk.Delete("/delete", produkController.Delete)
	produk.Get("/:id", produkController.Detail)

	// produk variasi
	produkVariasi := produk.Group("/variasi", middleware.Admin("marketing", "super admin"))
	produkVariasi.Post("/create", produkVariasiController.Create)
	produkVariasi.Put("/update", produkVariasiController.Update)
	produkVariasi.Delete("/delete", produkVariasiController.Delete)

	// promo
	promo := parent.Group("/promo", middleware.Admin("marketing", "super admin"))
	promo.Post("/create", promoController.Create)
	promo.Get("/list", promoController.Read)
	promo.Get("/:id", promoController.Detail)
	promo.Put("/update", promoController.Update)
	promo.Delete("/delete", promoController.Delete)

	// rekening
	rekening := parent.Group("/rekening", middleware.Admin("marketing", "super admin"))
	rekening.Post("/create", rekeningController.Create)
	rekening.Get("/list", rekeningController.Read)
	rekening.Put("/update", rekeningController.Update)
	rekening.Delete("/delete", rekeningController.Delete)

	// slider
	slider := parent.Group("/slider", middleware.Admin("marketing", "super admin"))
	slider.Post("/create", sliderController.Create)
	slider.Get("/list", sliderController.Read)
	slider.Put("/update", sliderController.Update)
	slider.Delete("/delete", sliderController.Delete)

	// user administratif
	userAdminstratif := parent.Group("/user-administratif", middleware.Admin("super admin"))
	userAdminstratif.Post("/create", userAdminController.Create)
	userAdminstratif.Get("/read", userAdminController.Read)
	userAdminstratif.Put("/update", userAdminController.Update)
	userAdminstratif.Delete("/delete", userAdminController.Delete)
}
