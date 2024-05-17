package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primary_key"`
	Name     string
	Products []Product `gorm:"many2many:product_categories;"`
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
	Categories []Category `gorm:"many2many:product_categories;"`
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
		Categories: []Category{category},
	})
	db.Create(&Product{
		Name:       "Book",
		Price:      10,
		Categories: []Category{category2},
	})

}
