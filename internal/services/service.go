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

type ProductService interface {
	CreateProduct(product models.Product, token string) error
	UpdateProduct(product models.Product, token string, id int) error
}
type Service struct {
	AuthService
	ProductService
}

func NewService(repo *repository.Repository) *Service {

	return &Service{
		AuthService:    newAuthService(repo),
		ProductService: newProductService(repo),
	}
}
