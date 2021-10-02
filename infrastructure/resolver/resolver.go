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
}

func MYSQLResolver(mysql *sql.DB) ServiceResolver {
	//setup repository
	adminRepo := repository.AdminRepository(mysql)
	kategoriRepo := repository.KategoriRepository(mysql)
	mitraRepo := repository.MitraRepository(mysql)

	//setup service
	adminServ := services.AdminService(&adminRepo)
	kategoriServ := services.KategoriService(&kategoriRepo)
	mitraServ := services.MitraService(&mitraRepo)

	return ServiceResolver{
		AdminService:    &adminServ,
		KategoriService: &kategoriServ,
		MitraService:    &mitraServ,
	}
}
