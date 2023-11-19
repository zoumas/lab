package main

import (
	"database/sql"
)

type Product struct {
	Name      string
	Price     float64
	Available bool
}

func DropProductTable(db *sql.DB) error {
	query := `DROP TABLE IF EXISTS products`
	_, err := db.Exec(query)
	return err
}

func CreateProductTable(db *sql.DB) error {
	query := `
  CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price NUMERIC(6, 2) NOT NULL,
    available BOOLEAN,
    created_at TIMESTAMP DEFAULT NOW()
  )`

	_, err := db.Exec(query)
	return err
}

func InsertProduct(db *sql.DB, product Product) (id int, err error) {
	query := `
  INSERT INTO products (name, price, available)
  VALUES ($1, $2, $3)
  RETURNING id`

	err = db.QueryRow(query, product.Name, product.Price, product.Available).Scan(&id)
	return
}

func GetAllProducts(db *sql.DB) ([]Product, error) {
	query := `SELECT name, price, available FROM products`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.Name, &product.Price, &product.Available); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
