package resolver

import (
	"database/sql"

	"github.com/smokers10/go-diagem.git/domain"
	"github.com/smokers10/go-diagem.git/repository"
	"github.com/smokers10/go-diagem.git/services"
)

type ServiceResolver struct {
	AdminService    *domain.AdminService
	KategoriService *domain.KategoriService
	MitraService    *domain.MitraService
	ProdukService   *domain.ProdukService
	PromoService    *domain.PromoService
	RekeningService *domain.RekeningService
}

func MYSQLResolver(mysql *sql.DB) ServiceResolver {
	//setup repository
	adminRepo := repository.AdminRepository(mysql)
	kategoriRepo := repository.KategoriRepository(mysql)
	mitraRepo := repository.MitraRepository(mysql)
	produkRepo := repository.ProdukRepository(mysql)
	promoRepo := repository.PromoRepository(mysql)
	rekeningRepo := repository.RekeningRepository(mysql)

	//setup service
	adminServ := services.AdminService(&adminRepo)
	kategoriServ := services.KategoriService(&kategoriRepo)
	mitraServ := services.MitraService(&mitraRepo)
	produkServ := services.ProdukService(&produkRepo)
	promoServ := services.PromoService(&promoRepo)
	rekeningServ := services.RekeningService(&rekeningRepo)

	return ServiceResolver{
		AdminService:    &adminServ,
		KategoriService: &kategoriServ,
		MitraService:    &mitraServ,
		ProdukService:   &produkServ,
		PromoService:    &promoServ,
		RekeningService: &rekeningServ,
	}
}
