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
	userAdministratifController := admin.UserAdministratifController()
	userController := admin.UserController()
	alamatOriginController := admin.AlamatOriginController()

	// controller API ini
	mitraAPIController := adminAPI.MitraController(resolver.MitraService)
	produkAPIController := adminAPI.ProdukController(resolver.ProdukService)
	produkFotoAPIController := adminAPI.ProdukFotoController(resolver.ProdukFoto)
	kategoriAPIController := adminAPI.KategoriController(resolver.KategoriService)
	varianAPIController := adminAPI.ProdukVariasiController(resolver.ProdukVariasiService)
	promoAPIController := adminAPI.PromoController(resolver.PromoService)
	sliderAPIController := adminAPI.SliderController(resolver.SliderService)
	userAdminController := adminAPI.UserAdminController(resolver.AdminService)
	userAPIController := adminAPI.UserController(resolver.UserService)
	alamatOriginAPIController := adminAPI.AlamatOriginController(resolver.AlamatOriginService)
	bankAPIController := adminAPI.BankController(resolver.BankService)
	rajaongkir := adminAPI.RajaOngkir(resolver.AlamatService)
	orderAPIController := adminAPI.OrderController(resolver.OrderService)

	// router clustering
	adminParentPath := app.Group("/admin")

	// authenctication
	adminParentPath.Get("/login", guestOnly, authenticationController.LoginPage)
	adminParentPath.Post("/login", guestOnly, authenticationController.Login)
	adminParentPath.Get("/logout", authenticationController.Logout)

	// reseller
	mitra := adminParentPath.Group("/reseller", middleware.AdminWeb(session, "admin", "super admin"))
	// reseller page
	mitra.Get("/", mitraController.IndexPage)
	mitra.Get("/tambah", mitraController.TambahPage)
	mitra.Get("/edit/:id", mitraController.EditPage)
	// reseller - action
	mitra.Get("/get", mitraAPIController.Read)
	mitra.Get("/get/:id", mitraAPIController.GetOne)
	mitra.Post("/create", mitraAPIController.Create)
	mitra.Put("/update", mitraAPIController.Update)
	mitra.Delete("/delete", mitraAPIController.Delete)

	// slider
	slider := adminParentPath.Group("/slider", middleware.AdminWeb(session, "admin", "super admin"))
	// slider page
	slider.Get("/", sliderController.IndexPage)
	slider.Get("/tambah", sliderController.TambahPage)
	slider.Get("/edit/:id", sliderController.EditPage)

	// slider - Action
	slider.Get("/get", sliderAPIController.Read)
	slider.Get("/get/:id", sliderAPIController.Detail)
	slider.Post("/create", sliderAPIController.Create)
	slider.Put("/update", sliderAPIController.Update)
	slider.Put("/update-cover", sliderAPIController.UpdateCover)
	slider.Delete("/delete", sliderAPIController.Delete)

	// keuangan
	keuangan := adminParentPath.Group("/keuangan", middleware.AdminWeb(session, "admin", "super admin"))
	keuangan.Get("/va", keuanganController.RekeningPage)
	keuangan.Put("/update", bankAPIController.Update)
	keuangan.Get("/get", bankAPIController.Read)

	// produk
	produk := adminParentPath.Group("/produk", middleware.AdminWeb(session, "admin", "super admin"))
	// produk - page
	produk.Get("/", produkController.IndexPage)
	produk.Get("/tambah", produkController.TambahPage)
	produk.Get("/edit/:id", produkController.EditPage)
	produk.Get("/kategori", produkController.KategoriPage)

	// produk - Action
	produk.Post("/create", produkAPIController.Create)
	produk.Get("/get", produkAPIController.Read)
	produk.Get("/get/:id", produkAPIController.Detail)
	produk.Delete("/delete", produkAPIController.Delete)
	produk.Put("/update", produkAPIController.Update)
	produk.Post("/foto/upload", produkFotoAPIController.Upload)
	produk.Put("/foto/foto-utama", produkFotoAPIController.MakeUtama)
	produk.Delete("/foto/delete", produkFotoAPIController.Delete)

	// varian - Action Only
	varian := adminParentPath.Group("/varian")
	varian.Post("/", varianAPIController.Create)
	varian.Put("/", varianAPIController.Update)
	varian.Delete("/", varianAPIController.Delete)

	// kategori
	kategori := adminParentPath.Group("/kategori", middleware.AdminWeb(session, "admin", "super admin"))
	// produk kategori - page
	kategori.Get("/", produkController.KategoriPage)
	kategori.Get("/tambah", produkController.KategoriTambahPage)
	kategori.Get("/edit/:id", produkController.KategoriEditPage)

	// produk kategori - action
	kategori.Get("/get", kategoriAPIController.Read)
	kategori.Get("/get/:id", kategoriAPIController.Detail)
	kategori.Post("/create", kategoriAPIController.Create)
	kategori.Put("/update", kategoriAPIController.Update)
	kategori.Put("/update-cover", kategoriAPIController.UpdateCover)
	kategori.Delete("/delete", kategoriAPIController.Delete)

	// order
	order := adminParentPath.Group("/order", middleware.AdminWeb(session, "admin", "super admin"))
	order.Get("/", orderController.IndexPage)
	order.Get("/get", orderAPIController.Read)
	order.Put("/update-state", orderAPIController.UpdateStatus)
	order.Get("/detail/:id", orderAPIController.DetailOrder)

	// promo
	promo := adminParentPath.Group("/promo", middleware.AdminWeb(session, "admin", "super admin"))
	// promo page
	promo.Get("/", promoController.IndexPage)
	promo.Get("/tambah", promoController.TambahPage)
	promo.Get("/edit/:id", promoController.EditPage)

	// promo - Action
	promo.Get("/get", promoAPIController.Read)
	promo.Get("/get/:id", promoAPIController.Detail)
	promo.Post("/create", promoAPIController.Create)
	promo.Put("/update", promoAPIController.Update)
	promo.Put("/update-cover", promoAPIController.UpdateCover)
	promo.Delete("/delete", promoAPIController.Delete)

	// user administratif
	userAdministratif := adminParentPath.Group("/user-administratif", middleware.AdminWeb(session, "admin", "super admin"))
	// user adminsitratif page
	userAdministratif.Get("/", userAdministratifController.IndexPage)
	userAdministratif.Get("/tambah", userAdministratifController.TambahPage)
	userAdministratif.Get("/edit/:id", userAdministratifController.EditPage)

	// user administratif - action
	userAdministratif.Get("/get", userAdminController.Read)
	userAdministratif.Get("/get/:id", userAdminController.Detail)
	userAdministratif.Post("/create", userAdminController.Create)
	userAdministratif.Put("/update", userAdminController.Update)
	userAdministratif.Put("/update-password", userAdminController.UpdatePassword)
	userAdministratif.Delete("/delete", userAdminController.Delete)

	// user
	user := adminParentPath.Group("/user", middleware.AdminWeb(session, "admin", "super admin", "marketing"))
	// user page
	user.Get("/", userController.IndexPage)
	user.Get("/profile/:id", userController.DetailPage)

	// user - action
	user.Get("/get", userAPIController.ReadAll)
	user.Get("/get/:id", userAPIController.Detail)

	// home
	adminParentPath.Get("/home", middleware.AdminWeb(session, "admin", "super admin", "marketing"), homeController.HomePage)

	// alamat origin
	alamatOrigin := adminParentPath.Group("/alamat-origin", middleware.AdminWeb(session, "admin", "super admin"))
	// alamat origin page
	alamatOrigin.Get("/", alamatOriginController.IndexPage)
	// alamat origin action
	alamatOrigin.Get("/read", alamatOriginAPIController.Read)
	alamatOrigin.Put("/update", alamatOriginAPIController.Update)

	// raja ongkir API
	rajaongkirPath := adminParentPath.Group("/rajaongkir", middleware.AdminWeb(session, "admin", "super admin"))
	rajaongkirPath.Get("/provinsi", rajaongkir.Provinsi)
	rajaongkirPath.Get("/kota/:provinsi_id", rajaongkir.Kota)
	rajaongkirPath.Get("/kota", rajaongkir.Kota)
}
