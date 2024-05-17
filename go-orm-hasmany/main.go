package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primary_key"`
	Name     string
	Products []Product
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
	category2 := Category{
		Name: "Books",
	}
	db.Create(&category2)
	db.Create(&Product{
		Name:       "Laptop",
		Price:      1000,
		CategoryID: category.ID,
	})
	db.Create(&Product{
		Name:       "Mobile",
		Price:      500,
		CategoryID: category2.ID,
	})
	db.Create(&Serial{
		Number:    "123456",
		ProductID: 1,
	})
	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.Serial").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name)
		for _, product := range category.Products {
			fmt.Println(product.Name, product.Price, product.Serial.Number)
		}
	}
}
