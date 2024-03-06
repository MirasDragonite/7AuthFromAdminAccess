package repository

import (
	"database/sql"
	"miras/internal/models"
)

type AuthRepo interface {
	CreateUser(user models.Register) error
	SelectUser(user models.Login) (models.User, error)
	GetSessionByUserID(id int64) (models.Session, error)
	GetSessionByToken(token string) (models.Session, error)
	CreateSession(session models.Session) error
	UpdateToken(token, date string, userID int) error
	GetUserById(id int) (models.User, error)
}
type ProductRepo interface {
	GetProduct(id int) (int, error)
	CreateProduct(product models.Product) error
	DeleteProduct(id int) error
	UpdateProduct(product models.Product) error
}

type Repository struct {
	AuthRepo
	ProductRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		AuthRepo:    newAuthRepo(db),
		ProductRepo: newProductRepo(db),
	}
}
