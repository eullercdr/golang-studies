package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primary_key"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// products := []Product{
	// 	{Name: "Mobile", Price: 20000},
	// 	{Name: "Tablet", Price: 10000},
	// }
	// db.Create(&products)

	var product Product
	db.Create(&Product{
		Name:  "Laptop",
		Price: 50000,
	})
	db.First(&product, 6)
	db.Delete(&product)
	fmt.Println(product)

	// db.First(&product, "name = ?", "Mobile")
	// fmt.Println(product)

	// var products []Product
	// db.Find(&products)
	// fmt.Println(products)

	// var products []Product
	// db.Offset(1).Limit(2).Find(&products)
	// fmt.Println(products)

	//where
	// var products []Product
	// db.Where("price > ?", 20).Find(&products)
	// fmt.Println(products)

	//like
	// var products []Product
	// db.Where("name LIKE ?", "%lap%").Find(&products)
	// fmt.Println(products)

	//delete
	// var p Product
	// db.First(&p, 1)
	// p.Name = "Laptop New"
	// db.Save(&p)

	// var p2 Product
	// db.First(&p2, 1)
	// fmt.Println(p2)

	// db.Delete(&p2)

}
