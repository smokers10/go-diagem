package resolver

import (
	"database/sql"

	"github.com/smokers10/go-diagem.git/domain"
	"github.com/smokers10/go-diagem.git/repository"
	"github.com/smokers10/go-diagem.git/services"
)

type ServiceResolver struct {
	AdminService        *domain.AdminService
	AlamatService       *domain.AlamatService
	KategoriService     *domain.KategoriService
	MitraService        *domain.MitraService
	ProdukService       *domain.ProdukService
	PromoService        *domain.PromoService
	RekeningService     *domain.RekeningService
	SliderService       *domain.SliderService
	UserService         *domain.UserService
	VerificationService *domain.VerifikasiService
}

func MYSQLResolver(mysql *sql.DB) ServiceResolver {
	//setup repository
	adminRepo := repository.AdminRepository(mysql)
	alamatRepo := repository.AlamatRepository(mysql)
	kategoriRepo := repository.KategoriRepository(mysql)
	mitraRepo := repository.MitraRepository(mysql)
	produkRepo := repository.ProdukRepository(mysql)
	promoRepo := repository.PromoRepository(mysql)
	rekeningRepo := repository.RekeningRepository(mysql)
	sliderRepo := repository.SliderRepository(mysql)
	userRepo := repository.UserRepository(mysql)
	verificationRepo := repository.VerificationRepository(mysql)

	//setup service
	adminServ := services.AdminService(&adminRepo)
	alamatServ := services.AlamtService(&alamatRepo)
	kategoriServ := services.KategoriService(&kategoriRepo)
	mitraServ := services.MitraService(&mitraRepo)
	produkServ := services.ProdukService(&produkRepo)
	promoServ := services.PromoService(&promoRepo)
	rekeningServ := services.RekeningService(&rekeningRepo)
	sliderServ := services.SliderService(&sliderRepo)
	userServ := services.UserService(&userRepo)
	verificationServ := services.VerificationService(&verificationRepo, &userRepo)

	return ServiceResolver{
		AdminService:        &adminServ,
		AlamatService:       &alamatServ,
		KategoriService:     &kategoriServ,
		MitraService:        &mitraServ,
		ProdukService:       &produkServ,
		PromoService:        &promoServ,
		RekeningService:     &rekeningServ,
		SliderService:       &sliderServ,
		UserService:         &userServ,
		VerificationService: &verificationServ,
	}
}
