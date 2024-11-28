package main

import (
	"github.com/eduardo1520/goexpert/9-APIs/internal/entity"
	"github.com/eduardo1520/goexpert/9-APIs/internal/infra/database"
	"github.com/eduardo1520/goexpert/9-APIs/internal/infra/webserver"
	"github.com/eduardo1520/goexpert/9-APIs/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	userDB := database.NewUser(db)

	productHandler := handlers.NewProductHandler(productDB)
	userHandler := handlers.NewUserHandler(userDB)

	r := webserver.SetupRoutes(productHandler, userHandler)
	http.ListenAndServe(":8000", r)

}
