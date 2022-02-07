package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/smokers10/go-diagem.git/domain"
	"github.com/smokers10/go-diagem.git/infrastructure/config"
	"github.com/smokers10/go-diagem.git/infrastructure/encryption"
	"github.com/smokers10/go-diagem.git/infrastructure/jwt"
)

type userServiceImpl struct {
	userRepository domain.UserRepository
}

type chaptacheck struct {
	Success     bool   `json:"success"`
	ChallengeTS string `json:"challenge_ts"`
}

func UserService(user *domain.UserRepository) domain.UserService {
	return &userServiceImpl{userRepository: *user}
}

// khusus admin
func (u *userServiceImpl) Read() *domain.Response {
	res := domain.Response{}

	user, err := u.userRepository.ReadAll()
	if err != nil {
		fmt.Println(err)
		res.Message = "error terjadi saat mengambil data user"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = user
	res.Message = "pengambilan data user berhasil"
	res.Status = http.StatusOK
	res.Success = true
	return &res
}

func (u *userServiceImpl) Detail(id int) *domain.Response {
	res := domain.Response{}

	user, err := u.userRepository.Detail(id)
	if err != nil {
		fmt.Println(err)
		res.Message = "error terjadi saat mengambil data user"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = user
	res.Message = "pengambilan data user berhasil"
	res.Status = http.StatusOK
	res.Success = true
	return &res
}

// khusus user
func (u *userServiceImpl) Login(cred *domain.UserCredential) *domain.Response {
	res := domain.Response{}
	payload := jwt.Payload{}

	// verifikasi chapta
	cresult, err := u.verifyChapta(cred.RechaptaResponse)
	if err != nil {
		fmt.Println(err)
		return &domain.Response{
			Message: "Error saat check chapta",
			Status:  500,
		}
	}

	if !cresult.Success {
		return &domain.Response{
			Message: "Silahkan checklist chapta terlebih dahulu",
			Status:  http.StatusUnauthorized,
		}
	}

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

func (u *userServiceImpl) Update(req *domain.UserBasicData) *domain.Response {
	// deklarasi variable
	res := domain.Response{}

	// panggil repository
	user, err := u.userRepository.Update(req)
	if err != nil {
		fmt.Println(err)
		res.Message = "error terjadi saat update akun Anda"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = user
	res.Message = "akun Anda berhasil diupdate"
	res.Status = http.StatusOK
	res.Success = true
	return &res
}

func (u *userServiceImpl) GetProfile(userID int) *domain.Response {
	// deklarasi variable
	res := domain.Response{}

	// panggil repository
	user, err := u.userRepository.ByID(userID)
	if err != nil {
		fmt.Println(err)
		res.Message = "error terjadi saat pengambilan akun"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = user
	res.Message = "akun berhasil diambil"
	res.Status = http.StatusOK
	res.Success = true
	return &res
}

func (a *userServiceImpl) verifyChapta(chaptaResponse string) (*chaptacheck, error) {
	// deklarasi var penting
	r := &chaptacheck{}
	c := config.ReadConfig().ETC
	u := "https://www.google.com/recaptcha/api/siteverify"

	// request HTTP dimulai
	resp, err := http.PostForm(u, url.Values{"secret": {c.RechaptaServerKey}, "response": {chaptaResponse}})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read body ke penyimpanan sementara (RAM)
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// unmarshal response body dari json ke struct
	json.Unmarshal(resBody, r)

	return r, nil
}
