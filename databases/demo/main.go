package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatalf("Error loading .env file : %q", err)
	}

	dsn, ok := env["DATA_SOURCE_NAME"]
	if !ok {
		log.Fatal("DATA_SOURCE_NAME is not set in .env")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database : %q", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	version, err := GetDatabaseVersion(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(version)

	if err := DropProductTable(db); err != nil {
		log.Fatal(err)
	}

	if err := CreateProductTable(db); err != nil {
		log.Fatal(err)
	}

	id, err := InsertProduct(db, Product{
		Name:      "Book",
		Price:     15.55,
		Available: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	var product Product
	row := db.QueryRow(`SELECT name, price, available FROM products WHERE id = $1`, id)
	if err := row.Scan(&product.Name, &product.Price, &product.Available); err != nil {
		log.Fatal(err)
	}
	fmt.Println(product)

	_, err = InsertProduct(db, Product{
		Name:      "Coffee",
		Price:     5.99,
		Available: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	products, err := GetAllProducts(db)
	if err != nil {
		log.Fatal(err)
	}
	for _, product := range products {
		fmt.Println(product)
	}
}

func GetDatabaseVersion(db *sql.DB) (string, error) {
	row := db.QueryRow(`SELECT version();`)
	if err := row.Err(); err != nil {
		return "", nil
	}

	var version string
	err := row.Scan(&version)
	return version, err
}
