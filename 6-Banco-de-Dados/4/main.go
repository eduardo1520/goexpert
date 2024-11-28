package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primarykey"`
	Name string
}

type Product struct {
	ID           int `gorm:primarykey`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primarykey"`
	Number    string
	ProductID int `gorm:"unique"` // Garante que o ProductID seja único
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Category{}, &Product{}, &SerialNumber{})

	// create category

	//category := Category{Name: "Eletronicos"}
	//db.Create(&category)

	//create product
	//var products = []Product{
	//	{Name: "Notebook", Price: 5000.00, CategoryID: category.ID},
	//	{Name: "Mouse", Price: 50.00, CategoryID: category.ID},
	//	{Name: "Keyboard", Price: 100.00, CategoryID: category.ID},
	//}
	//db.Create(products)

	//create serial number
	var serialNumbers = []SerialNumber{
		{Number: "123456", ProductID: 1},
		{Number: "123457", ProductID: 2},
		{Number: "123458", ProductID: 3},
	}
	db.Create(serialNumbers)

	var products []Product
	// Relation Has One (Um para um)
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}
}
