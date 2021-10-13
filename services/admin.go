package services

import (
	"fmt"
	"net/http"

	"github.com/smokers10/go-diagem.git/domain"
	"github.com/smokers10/go-diagem.git/infrastructure/encryption"
	"github.com/smokers10/go-diagem.git/infrastructure/jwt"
)

type adminServiceImpl struct {
	adminRepository domain.AdminRepository
}

func AdminService(admin *domain.AdminRepository) domain.AdminService {
	return &adminServiceImpl{adminRepository: *admin}
}

func (as *adminServiceImpl) Login(cred *domain.AdminCredentials) (response *domain.Response) {
	// deklarasi
	res := domain.Response{}
	payload := jwt.Payload{}

	// cari admin dengan email terkait
	admin, err := as.adminRepository.GetByEmail(cred.Email)
	if err != nil {
		fmt.Println(err.Error())
		res.Message = "terjadi error saat pencarian akun"
		res.Status = http.StatusInternalServerError
		return response
	}

	if admin.ID == 0 {
		res.Message = fmt.Sprintf("akun dengan email %s tidak ada", cred.Email)
		res.Status = http.StatusUnauthorized
		return &res
	}

	// check password
	if !encryption.Compare(admin.Password, cred.Password) {
		res.Message = "password salah"
		res.Status = http.StatusUnauthorized
		return &res
	}

	// buat token
	payload.Email = admin.Email
	payload.ID = admin.ID
	payload.Type = "admin"
	payload.IsVerified = false

	res.Message = "login berhasil"
	res.Data = admin
	res.Status = http.StatusOK
	res.Success = true
	res.Token = jwt.Sign(&payload)
	return &res
}
