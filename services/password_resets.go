package services

import (
	"github.com/smokers10/go-diagem.git/domain"
)

type passwordResetServiceImpl struct {
	passwordResetRepository domain.PasswordResetsRepository
	adminRepository         domain.AdminRepository
	userRepository          domain.UserRepository
}

// func PasswordResetsService(resetPassword *domain.PasswordResetsRepository, admin *domain.AdminRepository, user *domain.UserRepository) domain.PasswordResetsService {
// 	return &passwordResetServiceImpl{
// 		passwordResetRepository: *resetPassword,
// 		adminRepository:         *admin,
// 		userRepository:          *user,
// 	}
// }
