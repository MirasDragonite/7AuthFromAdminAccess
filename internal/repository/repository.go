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
}

type Repository struct {
	AuthRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{AuthRepo: newAuthRepo(db)}
}
