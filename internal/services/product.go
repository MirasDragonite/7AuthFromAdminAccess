package services

import (
	"errors"
	"miras/internal/models"
	"miras/internal/repository"
	"strings"
)

type Product struct {
	repo *repository.Repository
}

func newProductService(repo *repository.Repository) *Product {
	return &Product{repo: repo}
}
func (s *Product) CreateProduct(product models.Product, token string) error {

	session, err := s.repo.AuthRepo.GetSessionByToken(token)
	if err != nil {
		return err
	}

	user, err := s.repo.AuthRepo.GetUserById(session.UserID)
	if err != nil {
		return err
	}

	if user.Role != "admin" {
		return errors.New("access denied")
	}
	if checkIsThereEmptyField(product.Name, product.Category, product.ProductType, product.Year, product.AgeCategory, product.Chronology, product.KeyWords, product.Description, product.Director, product.Producer) {
		return errors.New("one of the fields is empty")
	}
	err = s.repo.ProductRepo.CreateProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *Product) UpdateProduct(product models.Product, token string, id int) error {

	session, err := s.repo.AuthRepo.GetSessionByToken(token)
	if err != nil {
		return err
	}

	user, err := s.repo.AuthRepo.GetUserById(session.UserID)
	if err != nil {
		return err
	}

	if user.Role != "admin" {
		return errors.New("access denied")
	}
	if checkIsThereEmptyField(product.Name, product.Category, product.ProductType, product.Year, product.AgeCategory, product.Chronology, product.KeyWords, product.Description, product.Director, product.Producer) {
		return errors.New("one of the fields is empty")
	}
	productID, err := s.repo.ProductRepo.GetProduct(id)
	if err != nil || productID < 1 {
		return errors.New("there is no product like this")
	}
	product.ID = int64(productID)
	err = s.repo.ProductRepo.UpdateProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func checkIsThereEmptyField(name, category, productType, year, ageCategory, chronology, keyWords, description, director, producer string) bool {

	if strings.TrimSpace(name) == "" || strings.TrimSpace(category) == "" || strings.TrimSpace(productType) == "" || strings.TrimSpace(year) == "" || strings.TrimSpace(ageCategory) == "" || strings.TrimSpace(chronology) == "" || strings.TrimSpace(keyWords) == "" || strings.TrimSpace(description) == "" || strings.TrimSpace(director) == "" || strings.TrimSpace(producer) == "" {
		return true
	}
	return false
}
