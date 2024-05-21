package main

import (
	"net/http"

	configs "github.com/eullercdr/go/api"
	"github.com/eullercdr/go/api/internal/entity"
	"github.com/eullercdr/go/api/internal/infra/database"
	"github.com/eullercdr/go/api/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	userDb := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDb, configs.TokenAuth, configs.JWTExpiresIn)
	productDb := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDb)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Get("/products", productHandler.GetProducts)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)

	http.ListenAndServe(":8080", r)
}
