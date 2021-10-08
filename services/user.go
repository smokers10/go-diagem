package services

import (
	"fmt"
	"net/http"

	"github.com/smokers10/go-diagem.git/domain"
	"github.com/smokers10/go-diagem.git/infrastructure/encryption"
	"github.com/smokers10/go-diagem.git/infrastructure/jwt"
)

type userServiceImpl struct {
	userRepository domain.UserRepository
}

func UserService(user *domain.UserRepository) domain.UserService {
	return &userServiceImpl{userRepository: *user}
}

func (u *userServiceImpl) Login(cred *domain.UserCredential) *domain.Response {
	res := domain.Response{}
	payload := jwt.Payload{}

	// check user
	user, err := u.userRepository.ByEmail(cred.Email)
	if err != nil {
		fmt.Println(err)
		res.Message = "error terjadi saat check ketersedian user"
		res.Status = http.StatusInternalServerError
		return &res
	}

	if user.ID == 0 {
		res.Message = fmt.Sprintf("pengguna dengan email %s tidak terdaftar", cred.Email)
		res.Status = http.StatusUnauthorized
		return &res
	}

	// check password
	if !encryption.Compare(user.Password, cred.Password) {
		res.Message = "password Anda salah"
		res.Status = http.StatusUnauthorized
		return &res
	}

	// buat payload token
	payload.Email = user.Email
	payload.ID = user.ID
	payload.Type = "user"
	payload.IsVerified = user.IsVerified

	// kosongkan passwword
	user.Password = ""

	// write response berhasil
	res.Data = user
	res.Message = "login berhasil"
	res.Status = http.StatusOK
	res.Success = true
	res.Token = jwt.Sign(&payload)

	return &res
}

func (u *userServiceImpl) Registrasi(req *domain.UserBasicData) *domain.Response {
	res := domain.Response{}
	payload := jwt.Payload{}

	// check user apakah terdaftar atau belum
	checkUser, err := u.userRepository.ByEmail(req.Email)
	if err != nil {
		fmt.Println(err)
		res.Message = "error terjadi saat check ketersedian user"
		res.Status = http.StatusInternalServerError
		return &res
	}

	if checkUser.ID != 0 {
		res.Message = fmt.Sprintf("pengguna dengan email %s sudah terdaftar", req.Email)
		res.Status = http.StatusConflict
		return &res
	}

	// encrypt password
	req.Password = encryption.Hash(req.Password)

	// simpan user baru
	newUser, err := u.userRepository.Create(req)
	if err != nil {
		fmt.Println(err)
		res.Message = "error terjadi saat menyimpan pengguna"
		res.Status = http.StatusInternalServerError
		return &res
	}

	// kosong kan data password
	newUser.Password = ""

	// buat payload JWT token
	payload.Email = newUser.Email
	payload.ID = newUser.ID
	payload.Type = "user"

	// write response berhasil
	res.Message = "registrasi berhasil"
	res.Status = http.StatusOK
	res.Success = true
	res.Data = newUser
	res.Token = jwt.Sign(&payload)

	return &res
}
