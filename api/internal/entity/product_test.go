package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Product 1", 100)
	assert.NotNil(t, product.ID)
	assert.Nil(t, err)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 100, product.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 100)
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("Product 1", 0)
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	product, err := NewProduct("Product 1", -1)
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidPrice, err)
}
