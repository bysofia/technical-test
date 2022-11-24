package routes

import (
	"nutech/handlers"
	"nutech/pkg/middleware"
	"nutech/pkg/postgre"
	"nutech/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryAuth(postgre.DB)
	h := handlers.HandlerAuth(userRepository)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST") // add this code
	r.HandleFunc("/check-auth", middleware.Auth(h.CheckAuth)).Methods("GET")
}