package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primarykey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:primarykey`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primarykey"`
	Number    string
	ProductID int
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

	//category := Category{Name: "Cozinha"}
	//db.Create(&category)

	//create product
	//db.Create(&Product{
	//	Name:       "Mouse",
	//	Price:      200.0,
	//	CategoryID: 1,
	//})

	//db.Create(&Product{
	//	Name:       "Panela",
	//	Price:      99.0,
	//	CategoryID: 1,
	//})
	//
	//db.Create(&SerialNumber{
	//	Number:    "123456",
	//	ProductID: 1,
	//})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	// Relation Has Many - (Um para muitos)
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("- ", product.Name, "Serial Number:", product.SerialNumber.Number)
		}
	}
}
