package domain

type PasswordResets struct {
	Email     string
	Token     string
	CreatedAt string
}

type PasswordResetsService interface {
}

type PasswordResetsRepository interface {
}
