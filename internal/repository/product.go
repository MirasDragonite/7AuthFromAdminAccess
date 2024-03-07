package repository

import (
	"database/sql"
	"miras/internal/models"
)

type Product struct {
	db *sql.DB
}

func newProductRepo(db *sql.DB) *Product {

	return &Product{db: db}
}

func (r *Product) GetProduct(id int) (int, error) {
	var result int
	query := `SELECT id FROM products WHERE id=$1`

	row := r.db.QueryRow(query, id)

	err := row.Scan(&result)
	if err != nil {

		return 0, err
	}

	return result, nil
}

func (r *Product) CreateProduct(product models.Product) error {

	query := `INSERT INTO products(name,category,product_type,year,age_category,chrono,key_words,description,director,producer) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`

	_, err := r.db.Exec(query, product.Name, product.Category, product.ProductType, product.Year, product.AgeCategory, product.Chronology, product.KeyWords, product.Description, product.Director, product.Producer)
	if err != nil {
		return err
	}

	return nil
}
func (r *Product) DeleteProduct(id int) error {
	query := `DELETE FROM products WHERE id=$1`

	_, err := r.db.Exec(query, &id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Product) UpdateProduct(product models.Product) error {

	query := `UPDATE products SET name=$1,category=$2,product_type=$3,year=$4,age_category=$5,chrono=$6,key_words=$7,description=$8,director=$9,producer=$10 WHERE id=$11`

	_, err := r.db.Exec(query, product.Name, product.Category, product.ProductType, product.Year, product.AgeCategory, product.Chronology, product.KeyWords, product.Description, product.Director, product.Producer, product.ID)
	if err != nil {
		return err
	}
	return nil
}
