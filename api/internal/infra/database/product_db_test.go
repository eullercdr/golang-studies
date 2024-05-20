package database

import (
	"fmt"
	"testing"

	"github.com/eullercdr/go/api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Errorf("Failed to open database: %v", err)
	}
	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Bicicleta", 100)
	productDb := NewProduct(db)

	err = productDb.Create(product)
	assert.Nil(t, err)

	var productFound entity.Product
	err = db.First(&productFound, "id=?", product.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
	assert.NotNil(t, productFound.ID)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Errorf("Failed to open database: %v", err)
	}
	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Bicicleta", 100)
	productDb := NewProduct(db)

	err = productDb.Create(product)
	assert.Nil(t, err)

	product.Name = "Bicicleta Elétrica Update"
	err = productDb.Update(product)
	assert.Nil(t, err)

	var productFound entity.Product
	err = db.First(&productFound, "id=?", product.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, "Bicicleta Elétrica Update", productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
	assert.NotNil(t, productFound.ID)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Errorf("Failed to open database: %v", err)
	}
	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Bicicleta", 100)
	productDb := NewProduct(db)

	err = productDb.Create(product)
	assert.Nil(t, err)

	err = productDb.Delete(product.ID.String())
	assert.Nil(t, err)

	var productFound entity.Product
	err = db.First(&productFound, "id=?", product.ID).Error
	assert.NotNil(t, err)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Errorf("Failed to open database: %v", err)
	}
	db.AutoMigrate(&entity.Product{})
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprint("Product ", i), 100)
		if err != nil {
			t.Errorf("Failed to create product: %v", err)
		}
		db.Create(product)
	}
	productDb := NewProduct(db)
	products, err := productDb.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.Equal(t, 10, len(products))
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDb.FindAll(2, 10, "asc")
	assert.Nil(t, err)
	assert.Equal(t, 10, len(products))
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)
}

func TestByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Errorf("Failed to open database: %v", err)
	}
	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Bicicleta", 100)
	productDb := NewProduct(db)

	err = productDb.Create(product)
	assert.Nil(t, err)

	productFound, err := productDb.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
	assert.NotNil(t, productFound.ID)
}
