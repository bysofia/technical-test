package routes

import (
	"nutech/handlers"
	"nutech/pkg/middleware"
	"nutech/pkg/mysql"
	"nutech/repositories"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
  productRepository := repositories.RepositoryProduct(mysql.DB)
  h := handlers.HandlerProduct(productRepository)

  r.HandleFunc("/product", middleware.Auth(middleware.UploadFile(h.CreateProduct))).Methods("POST")
  r.HandleFunc("/product/{id}", middleware.Auth(middleware.UploadFile(h.UpdateProduct))).Methods("PATCH")
  r.HandleFunc("/product/{id}", middleware.Auth(h.DeleteProduct)).Methods("DELETE")
  r.HandleFunc("/products", h.FindProducts).Methods("GET")
r.HandleFunc("/product/{id}", h.GetProduct).Methods("GET")
}