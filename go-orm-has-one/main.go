package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primary_key"`
	Name string
}

type Serial struct {
	ID        int `gorm:"primary_key"`
	Number    string
	ProductID int
}

type Product struct {
	ID         int `gorm:"primary_key"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category `gorm:"foreignKey:CategoryID"`
	Serial     Serial
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &Serial{})
	category := Category{
		Name: "Electronics",
	}
	db.Create(&category)
	db.Create(&Product{
		Name:       "Laptop",
		Price:      1000,
		CategoryID: category.ID,
	})
	db.Create(&Serial{
		Number:    "123456",
		ProductID: 1,
	})
	var products []Product
	db.Preload("Category").Preload("Serial").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.Serial.Number)
	}
}
