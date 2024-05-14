package main

import (
	"database/sql"

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
