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
	middlewareNotSoStrict := middleware.UserNotSoStrictWeb(session)

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
	postController := user.PostController(resolver.BlogService)
	pesananController := user.OrderController()
	resellerController := user.ResellerController()
	resetPasswordController := user.ResetPasswordController(resolver.ResetPasswordService)

	// API Controller init
	verificationAPIController := userAPI.VerificationController(resolver.VerificationService)
	alamatAPIController := userAPI.AlamatController(resolver.AlamatService)
	profileAPIController := userAPI.ProfileController(resolver.UserService)
	produkAPIController := userAPI.ProdukController(resolver.ProdukService)
	cartAPIController := userAPI.CartController(resolver.CartService)
	checkoutAPIController := userAPI.CheckoutController(resolver.CheckoutService)
	alamatOriginAPIController := userAPI.AlamatOriginController(resolver.AlamatOriginService)
	orderAPIController := userAPI.OrderController(resolver.OrderService)
	feedbackAPIController := userAPI.FeedbackController(resolver.FeedbackService)
	kategoriAPIController := userAPI.KategoriController(resolver.KategoriService)
	rajaongkirAPIController := userAPI.RajaOngkir(resolver.AlamatService)
	sellerAPIController := userAPI.MitraController(resolver.MitraService)

	// router clustering
	parentPath := app.Group("/")

	// authentication
	parentPath.Get("/login", middlewareGuestOnly, authencticationController.LoginPage)
	parentPath.Post("/login", middlewareGuestOnly, authencticationController.Login)
	parentPath.Get("/register", middlewareGuestOnly, authencticationController.RegisterPage)
	parentPath.Post("/register", middlewareGuestOnly, authencticationController.Register)
	parentPath.Get("/logout", authencticationController.Logout)

	resetPassword := parentPath.Group("forgot-password")
	resetPassword.Get("/", resetPasswordController.ForgotPasswordPage)
	resetPassword.Get("/reset/:token", resetPasswordController.ResetPasswordPage)
	resetPassword.Post("/create", resetPasswordController.Create)
	resetPassword.Post("/reset", resetPasswordController.Reset)

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
	alamat.Delete("/delete", alamatAPIController.Delete)
	alamat.Put("/make-utama", alamatAPIController.MakeUtama)
	alamat.Get("/provinsi", alamatAPIController.Provinsi)
	alamat.Get("/kota/:provinsi_id", alamatAPIController.Kota)

	// profile
	profile := parentPath.Group("/profile", middlewareStrict)
	profile.Get("/", profileController.ProfilePage)
	profile.Get("/get", profileAPIController.GetProfile)
	profile.Put("/update", profileAPIController.Update)

	// cart
	cart := parentPath.Group("/cart", middlewareStrict)
	cart.Get("/", cartController.IndexPage)
	cart.Get("/checkout", cartController.CheckoutPage)

	cart.Get("/read", cartAPIController.Read)
	cart.Post("/add-to-cart", cartAPIController.AddToCart)
	cart.Put("/update-quantity", cartAPIController.UpdateQuantity)
	cart.Delete("/delete", cartAPIController.DeleteCart)
	cart.Get("/total-carts", cartAPIController.CartCount)
	cart.Post("/ongkir", checkoutAPIController.Ongkir)
	cart.Post("/order", orderAPIController.Order)

	// produk
	produk := parentPath.Group("/produk", middlewareNotSoStrict)
	produk.Get("/", produkController.IndexPage)
	produk.Get("/detail/:slug", produkController.DetailPage)
	produk.Post("/get", produkAPIController.Read)
	produk.Get("/get/:slug", produkAPIController.Detail)
	produk.Get("/popular", produkAPIController.GetPopular)
	produk.Get("/kategori", kategoriAPIController.Read)

	// promo
	promo := parentPath.Group("/promo")
	promo.Get("/", promoController.IndexPage)
	promo.Get("/detail", promoController.DetailPage)

	// mitra
	mitra := parentPath.Group("/mitra")
	mitra.Get("/", mitraController.IndexPage)

	// post
	post := parentPath.Group("/blog", middlewareNotSoStrict)
	post.Get("/", postController.IndexPage)
	post.Get("/read/:slug", postController.ReadBlogPage)
	post.Get("/get", postController.GetAll)
	post.Get("/get/:slug", postController.ReadBlog)
	post.Get("/latest", postController.GetLatest)

	// pesanan page
	pesanan := parentPath.Group("/pesanan", middlewareStrict)
	pesanan.Get("/", pesananController.IndexPage)
	pesanan.Get("/invoice/:order_id", pesananController.Invoice)
	pesanan.Get("/get", orderAPIController.Read)
	pesanan.Get("/get/:order_id", orderAPIController.DetailOrder)

	// feedback
	feedback := parentPath.Group("/feedback")
	feedback.Get("/:produk_id", feedbackAPIController.Read)
	feedback.Get("/:produk_id/:rating", feedbackAPIController.ByRating)
	feedback.Post("/create", middlewareStrict, feedbackAPIController.GiveFeedback)
	feedback.Put("/update", middlewareStrict, feedbackAPIController.EditFeedback)

	// alamat origin
	alamatOrigin := parentPath.Group("/alamat-origin")
	alamatOrigin.Get("/", alamatOriginAPIController.Read, middlewareNotSoStrict)

	// reseller
	reseller := parentPath.Group("/seller", middlewareNotSoStrict)
	reseller.Get("/", resellerController.IndexPage)
	reseller.Get("/get", sellerAPIController.Get)

	// raja ongkir API
	rajaongkirPath := parentPath.Group("/rajaongkir")
	rajaongkirPath.Get("/kota", rajaongkirAPIController.Kota, middlewareStrict)

	// index page
	parentPath.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home", nil)
	})
}
