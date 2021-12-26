package resolver

import (
	"database/sql"

	"github.com/smokers10/go-diagem.git/domain"
	"github.com/smokers10/go-diagem.git/repository"
	"github.com/smokers10/go-diagem.git/services"
)

type ServiceResolver struct {
	AdminService         *domain.AdminService
	AlamatService        *domain.AlamatService
	AlamatOriginService  *domain.AlamatOriginService
	CartService          *domain.CartService
	KategoriService      *domain.KategoriService
	MitraService         *domain.MitraService
	ProdukService        *domain.ProdukService
	ProdukFoto           *domain.ProdukFotoService
	ProdukVariasiService *domain.ProdukVariasiService
	PromoService         *domain.PromoService
	SliderService        *domain.SliderService
	UserService          *domain.UserService
	VerificationService  *domain.VerifikasiService
	CheckoutService      *domain.CheckoutService
	OrderService         *domain.OrderService
	FeedbackService      *domain.FeedbackService
	BankService          *domain.BankService
}

func MYSQLResolver(mysql *sql.DB) ServiceResolver {
	//setup repository
	adminRepo := repository.AdminRepository(mysql)
	alamatRepo := repository.AlamatRepository(mysql)
	alamatOriginRepo := repository.AlamatOriginRepository(mysql)
	cartRepo := repository.CartRepository(mysql)
	kategoriRepo := repository.KategoriRepository(mysql)
	mitraRepo := repository.MitraRepository(mysql)
	produkRepo := repository.ProdukRepository(mysql)
	produkFoto := repository.ProdukFotoRepository(mysql)
	produkVariasiRepo := repository.ProdukVariasiRepository(mysql)
	promoRepo := repository.PromoRepository(mysql)
	sliderRepo := repository.SliderRepository(mysql)
	userRepo := repository.UserRepository(mysql)
	verificationRepo := repository.VerificationRepository(mysql)
	orderItemRepo := repository.OrderItemRepository(mysql)
	orderRepo := repository.OrderRepository(mysql)
	feedbackRepo := repository.FeedbackRepository(mysql)
	bankRepo := repository.BankRepository(mysql)

	//setup service
	adminServ := services.AdminService(&adminRepo)
	alamatServ := services.AlamtService(&alamatRepo)
	alamatOriginServ := services.AlamatOriginService(&alamatOriginRepo)
	cartServ := services.CartService(&cartRepo, &produkRepo, &produkVariasiRepo)
	kategoriServ := services.KategoriService(&kategoriRepo)
	mitraServ := services.MitraService(&mitraRepo)
	produkServ := services.ProdukService(&produkRepo, &produkFoto, &produkVariasiRepo)
	produkFotoServ := services.ProdukFoto(&produkFoto)
	produkVariasiServ := services.ProdukVariasiService(&produkVariasiRepo)
	promoServ := services.PromoService(&promoRepo)
	sliderServ := services.SliderService(&sliderRepo)
	userServ := services.UserService(&userRepo)
	verificationServ := services.VerificationService(&verificationRepo, &userRepo)
	checkoutServ := services.CheckoutService(&cartRepo)
	orderServ := services.OrderService(&orderRepo, &produkRepo, &produkVariasiRepo, &cartRepo, &orderItemRepo)
	feedbackServ := services.FeedbackService(&feedbackRepo)
	bankServ := services.BankService(&bankRepo)

	return ServiceResolver{
		AdminService:         &adminServ,
		AlamatService:        &alamatServ,
		AlamatOriginService:  &alamatOriginServ,
		CartService:          &cartServ,
		KategoriService:      &kategoriServ,
		MitraService:         &mitraServ,
		ProdukService:        &produkServ,
		ProdukFoto:           &produkFotoServ,
		ProdukVariasiService: &produkVariasiServ,
		PromoService:         &promoServ,
		SliderService:        &sliderServ,
		UserService:          &userServ,
		VerificationService:  &verificationServ,
		CheckoutService:      &checkoutServ,
		OrderService:         &orderServ,
		FeedbackService:      &feedbackServ,
		BankService:          &bankServ,
	}
}
