package database

import (
	"fmt"
	"github.com/eduardo1520/goexpert/9-APIs/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, errProduct := entity.NewProduct("Product 1", 10.00)
	assert.Nil(t, errProduct)
	productDB := NewProduct(db)
	errProduct = productDB.Create(product)
	assert.Nil(t, errProduct)
	assert.NotEmpty(t, product.ID)
}

func TestAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	for i := 1; i < 24; i++ {
		product, errProduct := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, errProduct)
		db.Create(product)
	}

	productDB := NewProduct(db)
	products, errProduct := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, errProduct)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, errProduct = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, errProduct)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, errProduct = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, errProduct)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func TestFindProductByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, errProduct := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, errProduct)

	db.Create(product)

	productDB := NewProduct(db)
	product, errProduct = productDB.FindById(product.ID.String())
	assert.NoError(t, errProduct)
	assert.Equal(t, "Product 1", product.Name)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, errProduct := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, errProduct)
	db.Create(product)
	productDB := NewProduct(db)
	product.Name = "Product 2"
	errProduct = productDB.Update(product)
	assert.NoError(t, errProduct)
	product, errProduct = productDB.FindById(product.ID.String())
	assert.NoError(t, errProduct)
	assert.Equal(t, "Product 2", product.Name)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	product, errProduct := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, errProduct)
	db.Create(product)
	productDB := NewProduct(db)
	errProduct = productDB.Delete(product.ID.String())
	assert.NoError(t, errProduct)
}
