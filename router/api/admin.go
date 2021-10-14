package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smokers10/go-diagem.git/controller/admin"
	"github.com/smokers10/go-diagem.git/infrastructure/middleware"
	"github.com/smokers10/go-diagem.git/infrastructure/resolver"
)

func AdminAPI(app *fiber.App, resolver *resolver.ServiceResolver) {
	//middleware
	superMiddleware := middleware.Admin("super admin")

	//controller
	authentication := admin.AdminAuthentication(resolver.AdminService)
	kategoriController := admin.KategoriController(resolver.KategoriService)
	mitraController := admin.MitraController(resolver.MitraService)
	produkController := admin.ProdukController(resolver.ProdukService)
	promoController := admin.PromoController(resolver.PromoService)
	rekeningController := admin.RekeningController(resolver.RekeningService)
	sliderController := admin.SliderController(resolver.SliderService)
	userAdminController := admin.UserAdminController(resolver.AdminService)

	//parent path
	parent := app.Group("/api/admin")

	//auth
	auth := parent.Group("/auth")
	auth.Post("/login", authentication.Login)

	//kategori
	kategori := parent.Group("/kategori")
	kategori.Post("/create", superMiddleware, kategoriController.Create)
	kategori.Get("/list", superMiddleware, kategoriController.Read)
	kategori.Put("/update", superMiddleware, kategoriController.Update)
	kategori.Delete("/delete", superMiddleware, kategoriController.Delete)

	//mitra
	mitra := parent.Group("/mitra")
	mitra.Post("/create", superMiddleware, mitraController.Create)
	mitra.Get("/list", superMiddleware, mitraController.Read)
	mitra.Get("/:id", superMiddleware, mitraController.GetOne)
	mitra.Put("/update", superMiddleware, mitraController.Update)
	mitra.Delete("/delete", superMiddleware, mitraController.Delete)

	//produk
	produk := parent.Group("/produk")
	produk.Post("/create", superMiddleware, produkController.Create)
	produk.Get("/list", superMiddleware, produkController.Read)
	produk.Put("/update", superMiddleware, produkController.Update)
	produk.Delete("/delete", superMiddleware, produkController.Delete)
	produk.Get("/:id", superMiddleware, produkController.Detail)

	//promo
	promo := parent.Group("/promo")
	promo.Post("/create", superMiddleware, promoController.Create)
	promo.Get("/list", superMiddleware, promoController.Read)
	promo.Get("/:id", superMiddleware, promoController.Detail)
	promo.Put("/update", superMiddleware, promoController.Update)
	promo.Delete("/delete", superMiddleware, promoController.Delete)

	//rekening
	rekening := parent.Group("/rekening")
	rekening.Post("/create", superMiddleware, rekeningController.Create)
	rekening.Get("/list", superMiddleware, rekeningController.Read)
	rekening.Put("/update", superMiddleware, rekeningController.Update)
	rekening.Delete("/delete", superMiddleware, rekeningController.Delete)

	//slider
	slider := parent.Group("/slider")
	slider.Post("/create", superMiddleware, sliderController.Create)
	slider.Get("/list", superMiddleware, sliderController.Read)
	slider.Put("/update", superMiddleware, sliderController.Update)
	slider.Delete("/delete", superMiddleware, sliderController.Delete)

	//user administratif
	userAdminstratif := parent.Group("/user-administratif")
	userAdminstratif.Post("/create", superMiddleware, userAdminController.Create)
	userAdminstratif.Get("/read", superMiddleware, userAdminController.Read)
	userAdminstratif.Put("/update", superMiddleware, userAdminController.Update)
	userAdminstratif.Delete("/delete", superMiddleware, userAdminController.Delete)

}
