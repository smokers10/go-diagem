package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/smokers10/go-diagem.git/domain"
	"github.com/smokers10/go-diagem.git/infrastructure/config"
	Email "github.com/smokers10/go-diagem.git/infrastructure/email"
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

	// kirim kode reset password via email
	redirect := config.ReadConfig().Application.APP_Base_URL + "/forgot-password" + "/reset/" + token.String()
	template := Email.SMTP().ForgotPasswordTemplate(user.Nama, kode, redirect)
	if err := Email.SMTP().NativeFire([]string{email}, "Lupa Password", template); err != nil {
		fmt.Println(err)
		return &domain.Response{
			Message: "error saat kirim email",
			Status:  500,
		}
	}

	return &domain.Response{
		Message: "Request Berhasil",
		Success: true,
		Status:  200,
	}
}

func (prs *passwordResetServiceImpl) Reset(req *domain.PasswordResets) *domain.Response {
	// ambil data reset password berdasarkan token
	resetPWData, err := prs.passwordResetRepository.ByToken(req.Token)
	if err != nil {
		fmt.Println(err)
		return &domain.Response{
			Message: "error saat mengambil data reset password",
			Status:  500,
		}
	}

	// validasi keberadaan reset password
	if resetPWData.Token == "" {
		return &domain.Response{
			Message: "Sesi Reset Password Tidak Valid",
			Status:  200,
		}
	}

	// comparasi kode inputan dengan kode dari database
	if !encryption.Compare(resetPWData.Kode, req.Kode) {
		return &domain.Response{
			Message: "Kode Reset Password Tidak Valid",
			Status:  200,
		}
	}

	// ambil data user
	user, err := prs.userRepository.ByEmail(resetPWData.Email)
	if err != nil {
		return &domain.Response{
			Message: "error saat mengambil data user",
			Status:  500,
		}
	}

	// update password
	if err := prs.userRepository.UpdatePassword(encryption.Hash(req.NewPassword), user.ID); err != nil {
		fmt.Println(err)
		return &domain.Response{
			Message: "error saat mengupdate password",
			Status:  500,
		}
	}

	// hapus sesi reset password
	if err := prs.passwordResetRepository.Delete(resetPWData.Email); err != nil {
		return &domain.Response{
			Message: "error saat menghapus sesi reset password",
			Status:  500,
		}
	}

	// final response
	return &domain.Response{
		Message: "Password Berhasil Diubah",
		Success: true,
	}
}
