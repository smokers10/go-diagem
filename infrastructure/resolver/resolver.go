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
}

func MYSQLResolver(mysql *sql.DB) ServiceResolver {
	//setup repository
	adminRepo := repository.AdminRepository(mysql)
	kategoriRepo := repository.KategoriRepository(mysql)

	//setup service
	adminServ := services.AdminService(&adminRepo)
	kategoriServ := services.KategoriService(&kategoriRepo)

	return ServiceResolver{
		AdminService:    &adminServ,
		KategoriService: &kategoriServ,
	}
}
