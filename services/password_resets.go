package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/smokers10/go-diagem.git/domain"
	"github.com/smokers10/go-diagem.git/infrastructure/encryption"
	"github.com/smokers10/go-diagem.git/infrastructure/etc"
)

type passwordResetServiceImpl struct {
	passwordResetRepository domain.PasswordResetsRepository
	userRepository          domain.UserRepository
}

func PasswordResetService(password *domain.PasswordResetsRepository, user *domain.UserRepository) domain.PasswordResetsService {
	return &passwordResetServiceImpl{passwordResetRepository: *password, userRepository: *user}
}

func (prs *passwordResetServiceImpl) Create(email string) *domain.Response {
	// check ketersediaan user dengan email
	user, err := prs.userRepository.ByEmail(email)
	if err != nil {
		fmt.Println(err)
		return &domain.Response{
			Message: "error saat check email",
			Status:  500,
		}
	}

	if user.Email == "" {
		return &domain.Response{
			Message: "Request Gagal!",
		}
	}

	// check apakah request dengan email terkait sudah ada atau belum
	currentRPW, err := prs.passwordResetRepository.ByEmail(email)
	if err != nil {
		fmt.Println(err)
		return &domain.Response{
			Message: "error saat check reset password",
			Status:  500,
		}
	}

	// buat token
	token, _ := uuid.NewRandom()
	kode := etc.KodeGeneratorImproved(6)
	fmt.Println(kode)

	// simpan request reset
	req := &domain.PasswordResets{
		Email: email,
		Token: token.String(),
		Kode:  encryption.Hash(kode),
	}

	// jika currentRPW tidak ada buat baru
	if currentRPW.Token == "" {
		prs.passwordResetRepository.Create(req)
	}

	// jika currentRPW ada, update yang ada
	if currentRPW.Token != "" {
		prs.passwordResetRepository.Update(req)
	}

	return &domain.Response{
		Message: "Request Berhasil",
		Success: true,
		Status:  200,
	}
}

func (prs *passwordResetServiceImpl) Reset(req *domain.PasswordResets) *domain.Response {
	res := domain.Response{}

	return &res
}
