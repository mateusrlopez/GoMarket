package routes

import (
	"github.com/gorilla/mux"
	"github.com/mateusrlopez/go-market/handlers"
	"github.com/mateusrlopez/go-market/repositories"
	"gorm.io/gorm"
)

func setupAuthRoutes(r *mux.Router, db *gorm.DB) {
	rep := repositories.UserRepository{DB: db}
	h := handlers.AuthHandler{Repository: rep}

	r.HandleFunc("/auth/register", h.Register).Methods("POST")
	r.HandleFunc("/auth/login", h.Login).Methods("POST")
}
