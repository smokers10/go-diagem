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
	payload.Type = admin.Roles
	payload.IsVerified = false

	res.Message = "login berhasil"
	res.Data = admin
	res.Status = http.StatusOK
	res.Success = true
	res.Token = jwt.Sign(&payload)
	return &res
}

func (as *adminServiceImpl) Create(req *domain.Admin) *domain.Response {
	// deklarasi
	res := domain.Response{}

	// hash password
	req.Password = encryption.Hash(req.Password)

	// panggil repository
	admins, err := as.adminRepository.Create(req)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat menambah user administratif"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = admins
	res.Message = fmt.Sprintf("akun %s berhasil ditambahkan", req.Roles)
	res.Status = http.StatusOK
	res.Success = true
	return &res
}

func (as *adminServiceImpl) Update(req *domain.Admin, logged_id int, logged_role string) *domain.Response {
	// deklarasi variable
	res := domain.Response{}

	// check super admin
	checkAdmins, err := as.adminRepository.GetByID(req.ID)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat check user administratif"
		res.Status = http.StatusInternalServerError
		return &res
	}

	// hanya boleh update dirinya sendiri
	if (checkAdmins.ID == logged_id) || (logged_role == "super admin") {
		// panggil repository
		admins, err := as.adminRepository.Update(req)
		if err != nil {
			fmt.Println(err)
			res.Message = "error saat update user administratif"
			res.Status = http.StatusInternalServerError
			return &res
		}

		res.Data = admins
		res.Message = fmt.Sprintf("akun %s berhasil ditambahkan", req.Nama)
		res.Status = http.StatusOK
		res.Success = true
		return &res
	}

	res.Message = "akses tidak diizinkan"
	res.Status = http.StatusForbidden
	return &res
}

func (as *adminServiceImpl) Delete(id int) *domain.Response {
	// deklarasi variable
	res := domain.Response{}

	// panggil repository
	if err := as.adminRepository.Delete(id); err != nil {
		fmt.Println(err)
		res.Message = "error saat menghapus user administratif"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "akun administratif berhasil dihapus"
	res.Status = http.StatusOK
	res.Success = true
	return &res
}

func (as *adminServiceImpl) Read(logged_id int) *domain.Response {
	// deklarasi variable
	res := domain.Response{}

	// panggil repository
	admin, err := as.adminRepository.Read(logged_id)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil user administratif"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = admin
	res.Message = "akun administratif berhasil diambil"
	res.Status = http.StatusOK
	res.Success = true
	return &res
}
