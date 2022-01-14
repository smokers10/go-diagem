package services

import (
	"fmt"
	"net/http"

	"github.com/smokers10/go-diagem.git/domain"
	"github.com/smokers10/go-diagem.git/infrastructure/config"
	"github.com/smokers10/go-diagem.git/infrastructure/encryption"
	"github.com/smokers10/go-diagem.git/infrastructure/etc"
	"github.com/smokers10/go-diagem.git/infrastructure/jwt"
)

type verificationServiceImpl struct {
	verificationRepository domain.VerifikasiRepository
	userRepository         domain.UserRepository
}

func VerificationService(verification *domain.VerifikasiRepository, user *domain.UserRepository) domain.VerifikasiService {
	return &verificationServiceImpl{
		verificationRepository: *verification,
		userRepository:         *user,
	}
}

func (v *verificationServiceImpl) Create(req *domain.Verifikasi) *domain.Response {
	res := domain.Response{}

	// check apakah user_id sudah ada
	verification, err := v.verificationRepository.ByUserID(req.UserID)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat check verifikasi"
		res.Status = http.StatusInternalServerError
		return &res
	}

	// generate kode verifikasi
	kode := etc.KodeGeneratorImproved(6)

	// encrypt kode verifikasi
	req.Kode = encryption.Hash(kode)

	if mode := config.ReadConfig().Application.APP_Production_Mode; mode == "development" {
		fmt.Println(kode)
	}

	// jika verifikasi dengan user_id sudah ada update jika tidak buat baru.
	if verification.UserID != 0 {
		if err := v.verificationRepository.Update(req); err != nil {
			fmt.Println(err)
			res.Message = "error saat memperbaharui verifikasi"
			res.Status = http.StatusInternalServerError
			return &res
		}
	} else {
		if err := v.verificationRepository.Create(req); err != nil {
			fmt.Println(err)
			res.Message = "error saat membuat verifikasi"
			res.Status = http.StatusInternalServerError
			return &res
		}
	}

	// write response sukses
	res.Message = "kode verifikasi dikirim ke email Anda"
	res.Status = http.StatusOK
	res.Success = true
	return &res
}

func (v *verificationServiceImpl) Verificate(req *domain.Verifikasi) *domain.Response {
	res := domain.Response{}
	payload := jwt.Payload{}

	// ambil data verifikasi
	verif, err := v.verificationRepository.ByUserID(req.UserID)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil data verifikasi"
		res.Status = http.StatusInternalServerError
		return &res
	}

	// jika data verifikasi tidak ada
	if verif.UserID == 0 {
		res.Message = "tidak ada data verifikasi untuk akun anda"
		res.Status = http.StatusOK
		return &res
	}

	// komparasi kode
	if !encryption.Compare(verif.Kode, req.Kode) {
		res.Message = "kode verifikasi salah"
		res.Status = http.StatusOK
		return &res
	}

	// ubah status verifikasi pada user
	if err := v.userRepository.UpdateVerificationStatus(verif.UserID); err != nil {
		fmt.Println(err)
		res.Message = "error saat memverifikasi akun"
		res.Status = http.StatusInternalServerError
		return &res
	}

	// ambil data user yang baru saja di verifikasi
	user, err := v.userRepository.ByID(req.UserID)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil data akun"
		res.Status = http.StatusInternalServerError
		return &res
	}

	// hapus data verifikasi
	if err := v.verificationRepository.Delete(req.UserID); err != nil {
		fmt.Println(err)
		res.Message = "error saat revoke sesi verifikasi"
		res.Status = http.StatusInternalServerError
		return &res
	}

	// setup JWT payload
	payload.Email = user.Email
	payload.ID = user.ID
	payload.IsVerified = true
	payload.Type = "user"

	res.Message = "akun anda terverifikasi"
	res.Status = http.StatusOK
	res.Success = true
	res.Token = jwt.Sign(&payload)

	return &res
}
