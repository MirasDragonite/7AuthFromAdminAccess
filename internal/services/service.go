package services

import (
	"miras/internal/models"
	"miras/internal/repository"
	"net/http"
)

type AuthService interface {
	CreateUser(user models.Register) error
	LoginToSystem(login models.Login) (http.Cookie, error)
}
type Service struct {
	AuthService
}

func NewService(repo *repository.Repository) *Service {

	return &Service{AuthService: newAuthService(repo)}
}
