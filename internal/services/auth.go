package services

import (
	"database/sql"
	"miras/internal/models"
	"miras/internal/repository"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	repo *repository.Repository
}

func newAuthService(repo *repository.Repository) *Auth {
	return &Auth{repo: repo}
}
func (s *Auth) CreateUser(user models.Register) error {

	hashPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPassword
	return s.repo.AuthRepo.CreateUser(user)
}

func (s *Auth) LoginToSystem(login models.Login) (http.Cookie, error) {

	user, err := s.repo.AuthRepo.SelectUser(login)
	if err != nil {
		return http.Cookie{}, err
	}

	cookie := http.Cookie{
		Path:     "/",
		HttpOnly: true,
		Name:     "Token",
	}
	newToken, err := createToken(login.Email)
	if err != nil {
		return http.Cookie{}, err
	}
	session, err := s.repo.AuthRepo.GetSessionByUserID(user.ID)
	newTime := time.Now()

	if err != nil {
		if err == sql.ErrNoRows {

			session.ID = int(user.ID)
			session.Token = newToken
			session.ExpiredDate = newTime.Format("2006-01-02 15:04:05")
			cookie.Value = newToken
			cookie.Expires = newTime
			err = s.repo.AuthRepo.CreateSession(session)
			if err != nil {
				return http.Cookie{}, err
			}
			return cookie, nil

		} else {
			return http.Cookie{}, err
		}

	} else {

		sessionTime, err := time.Parse("2006-01-02 15:04:05", session.ExpiredDate)
		if err != nil {
			return http.Cookie{}, err
		}

		if sessionTime.Before(time.Now()) || sessionTime.Equal(time.Now()) {

			err = s.repo.AuthRepo.UpdateToken(newToken, newTime.Format("2006-01-02 15:04:05"), int(user.ID))
			if err != nil {
				return http.Cookie{}, nil
			}
			cookie.Value = newToken
			cookie.Expires = newTime

		} else {
			cookie.Value = session.Token
			cookie.Expires = sessionTime
		}
	}
	return cookie, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

var secretKey = []byte("secret-key")

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Minute * 30).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
