package webserver

import (
	"github.com/eduardo1520/goexpert/9-APIs/configs"
	"github.com/eduardo1520/goexpert/9-APIs/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
)

func SetupRoutes(productHandler *handlers.ProductHandler, userHandler *handlers.UserHandler) chi.Router {
	c, err := configs.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.WithValue("jwt", c.TokenAuth))
	r.Use(middleware.WithValue("experiesin", int64(c.JWtExperiesIn)))

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(c.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Get("/", productHandler.GetProducts)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.Create)
	r.Post("/users/generate_token", userHandler.GetJWT)

	return r

}
