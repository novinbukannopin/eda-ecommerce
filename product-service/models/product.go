package models

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CategoryId  int64   `json:"category_id"`
}

type ProductCategory struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
