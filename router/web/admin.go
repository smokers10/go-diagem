package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	adminAPI "github.com/smokers10/go-diagem.git/controller/api/admin"
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
	mitraController := admin.MitraController()
	sliderController := admin.SliderController()
	keuanganController := admin.KeuanganController()
	produkController := admin.ProdukController()
	orderController := admin.OrderController()
	promoController := admin.PromoController()

	// controller API ini
	mitraAPIController := adminAPI.MitraController(resolver.MitraService)
	produkAPIController := adminAPI.ProdukController(resolver.ProdukService)
	kategoriController := adminAPI.KategoriController(resolver.KategoriService)
	varianController := adminAPI.ProdukVariasiController(resolver.ProdukVariasiService)

	// router clustering
	adminParentPath := app.Group("/admin")

	// authenctication
	adminParentPath.Get("/login", guestOnly, authenticationController.LoginPage)
	adminParentPath.Post("/login", guestOnly, authenticationController.Login)
	adminParentPath.Get("/logout", authenticationController.Logout)

	// mitra
	mitra := adminParentPath.Group("/mitra", middleware.AdminWeb(session, "admin", "super admin"))
	// mitra page
	mitra.Get("/", mitraController.IndexPage)
	mitra.Get("/tambah", mitraController.TambahPage)
	mitra.Get("/edit", mitraController.EditPage)
	// mitra non page
	mitra.Get("/get", mitraAPIController.Read)

	// slider
	slider := adminParentPath.Group("/slider", middleware.AdminWeb(session, "admin", "super admin"))
	slider.Get("/", sliderController.IndexPage)
	slider.Get("/tambah", sliderController.TambahPage)
	slider.Get("/edit", sliderController.EditPage)

	// keuangan
	keuangan := adminParentPath.Group("/keuangan", middleware.AdminWeb(session, "admin", "super admin"))
	keuangan.Get("/rekening", keuanganController.RekeningPage)

	// produk
	produk := adminParentPath.Group("/produk", middleware.AdminWeb(session, "admin", "super admin"))
	produk.Get("/", produkController.IndexPage)
	produk.Get("/tambah", produkController.TambahPage)
	produk.Get("/edit/:id", produkController.EditPage)
	produk.Get("/kategori", produkController.KategoriPage)

	// produk - Action
	produk.Post("/tambah", produkAPIController.Create)
	produk.Get("/get", produkAPIController.Read)
	produk.Get("/get/:id", produkAPIController.Detail)
	produk.Delete("/delete", produkAPIController.Delete)

	// varian - Action Only
	varian := adminParentPath.Group("/varian")
	varian.Post("/", varianController.Create)
	varian.Put("/", varianController.Update)
	varian.Delete("/", varianController.Delete)

	// kategori
	kategori := adminParentPath.Group("/kategori")
	kategori.Get("/get", kategoriController.Read)

	// order
	order := adminParentPath.Group("/order", middleware.AdminWeb(session, "admin", "super admin"))
	order.Get("/", orderController.IndexPage)

	// promo
	promo := adminParentPath.Group("/promo", middleware.AdminWeb(session, "admin", "super admin"))
	promo.Get("/", promoController.IndexPage)
	promo.Get("/tambah", promoController.TambahPage)
	promo.Get("/edit", promoController.EditPage)

	// home
	adminParentPath.Get("/home", middleware.AdminWeb(session, "admin", "super admin", "marketing"), homeController.HomePage)
}
