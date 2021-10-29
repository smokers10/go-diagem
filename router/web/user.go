package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	userAPI "github.com/smokers10/go-diagem.git/controller/api/user"
	"github.com/smokers10/go-diagem.git/controller/web/user"

	"github.com/smokers10/go-diagem.git/infrastructure/middleware"
	"github.com/smokers10/go-diagem.git/infrastructure/resolver"
)

func UserWebPage(app *fiber.App, session *session.Store, resolver *resolver.ServiceResolver) {
	// middleware init
	middlewareStrict := middleware.UserStrictWeb(session)
	middlewareVerification := middleware.UserVerificationWeb(session)
	middlewareGuestOnly := middleware.GuestOnly(session)

	// controller init
	authencticationController := user.AuthenticationController(resolver.UserService, session)
	homeController := user.HomeController()
	verifikasiController := user.VerificationController(resolver.VerificationService, session)
	alamatController := user.AlamatController()
	profileController := user.UserProfileCOntroller()
	cartController := user.CartController()
	produkController := user.ProdukController()
	promoController := user.PromoController()
	mitraController := user.MitraController()
	postController := user.PostController()

	// API Controller init
	verificationAPIController := userAPI.VerificationController(resolver.VerificationService)
	alamatAPIController := userAPI.AlamatController(resolver.AlamatService)
	profileAPIController := userAPI.ProfileController(resolver.UserService)

	// router clustering
	parentPath := app.Group("/")

	// authentication
	parentPath.Get("/login", middlewareGuestOnly, authencticationController.LoginPage)
	parentPath.Post("/login", middlewareGuestOnly, authencticationController.Login)
	parentPath.Get("/register", middlewareGuestOnly, authencticationController.RegisterPage)
	parentPath.Post("/register", middlewareGuestOnly, authencticationController.Register)
	parentPath.Get("/logout", authencticationController.Logout)

	// verifikasi
	verifikasi := parentPath.Group("/verifikasi", middlewareVerification)
	verifikasi.Get("/", verifikasiController.VerificationPage)
	verifikasi.Post("/create", verificationAPIController.Create)
	verifikasi.Post("/verificate", verifikasiController.Verificate)

	// home
	parentPath.Get("/home", middlewareStrict, homeController.HomePage)

	// alamat
	alamat := parentPath.Group("/alamat", middlewareStrict)
	alamat.Get("/", alamatController.AlamatPage)
	alamat.Post("/create", alamatAPIController.Create)
	alamat.Get("/read", alamatAPIController.Read)
	alamat.Put("/update", alamatAPIController.Update)
	alamat.Delete("/Delete", alamatAPIController.Delete)

	// profile
	profile := parentPath.Group("/profile", middlewareStrict)
	profile.Get("/", profileController.ProfilePage)
	profile.Get("/get", profileAPIController.GetProfile)
	profile.Put("/update", profileAPIController.Update)

	// cart
	cart := parentPath.Group("/cart", middlewareStrict)
	cart.Get("/", cartController.IndexPage)
	cart.Get("/checkout", cartController.CheckoutPage)
	cart.Get("/pembayaran", cartController.PembayaranPage)

	// produk
	produk := parentPath.Group("/produk")
	produk.Get("/", produkController.IndexPage)
	produk.Get("/detail", produkController.DetailPage)
	produk.Get("/produk-detail", produkController.ProdukDetailPage)

	// promo
	promo := parentPath.Group("/promo")
	promo.Get("/", promoController.IndexPage)
	promo.Get("/detail", promoController.DetailPage)

	// mitra
	mitra := parentPath.Group("/mitra")
	mitra.Get("/", mitraController.IndexPage)

	// post
	post := parentPath.Group("/blog")
	post.Get("/", postController.IndexPage)

	// index page
	parentPath.Get("/", middlewareGuestOnly, func(c *fiber.Ctx) error {
		return c.Render("home", nil)
	})
}
