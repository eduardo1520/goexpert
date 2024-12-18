package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primarykey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID         int `gorm:primarykey`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Category{}, &Product{})

	//category := Category{Name: "Eletronicos"}
	//db.Create(&category)
	//
	//category2 := Category{Name: "Cozinha"}
	//db.Create(&category2)

	//create product
	//db.Create(&Product{
	//	Name:       "Panela",
	//	Price:      99.0,
	//	Categories: []Category{category, category2},
	//})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	// Relation Many to Many - (Muitos para muitos)
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("- ", product.Name)
		}
	}
}
