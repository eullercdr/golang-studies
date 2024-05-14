package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //underscore is used to import a package only for its side effects

	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	product := NewProduct("Laptop", 50000)
	err = InsertProduct(db, product)
	if err != nil {
		panic(err)
	}
	product.Price = 100
	err = UpdateProduct(db, product)
	if err != nil {
		panic(err)
	}
	p, err := findOneProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	println(p.Name, p.Price)
	err = RemoveProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	products, err := findAll(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("Product: %s, Price: %f\n", p.Name, p.Price)
	}
}

func InsertProduct(db *sql.DB, p *Product) error {
	stmt, err := db.Prepare("INSERT INTO products(id, name, price) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.ID, p.Name, p.Price) //_ is used to ignore the first return value that is result
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(db *sql.DB, p *Product) error {
	stmt, err := db.Prepare("UPDATE products SET name=?, price=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.Name, p.Price, p.ID)
	if err != nil {
		return err
	}
	return nil
}

func findOneProduct(db *sql.DB, id string) (*Product, error) {
	smt, err := db.Prepare("SELECT id, name, price FROM products WHERE id=?")
	if err != nil {
		return nil, err
	}
	defer smt.Close()
	var p Product
	err = smt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price) //Scan is used to scan the result of the query
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func findAll(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func RemoveProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM products WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
