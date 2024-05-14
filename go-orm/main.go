package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primary_key"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/gorm"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})
	// db.Create(&Product{
	// 	Name:  "Laptop",
	// 	Price: 50000,
	// })
	products := []Product{
		{Name: "Mobile", Price: 20000},
		{Name: "Tablet", Price: 10000},
	}
	db.Create(&products)
}
