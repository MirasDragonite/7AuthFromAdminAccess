package models

type Product struct {
	ID          int64
	Name        string `json:"name"`
	Category    string `json:"category"`
	ProductType string `json:"product_type"`
	Year        string `json:"year"`
	AgeCategory string `json:"age_category"`
	Chronology  string `json:"chronology"`
	KeyWords    string `json:"key_words"`
	Description string `json:"description"`
	Director    string `json:"director"`
	Producer    string `json:"producer"`
}
