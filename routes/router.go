package routes

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/mateusrlopez/go-market/database"
	"github.com/mateusrlopez/go-market/handlers"
	"github.com/mateusrlopez/go-market/middlewares"
	"github.com/mateusrlopez/go-market/repositories"
	"github.com/mateusrlopez/go-market/settings"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.Use(middlewares.HeadersMiddleware)

	sr := r.PathPrefix(fmt.Sprintf("/%s", settings.Settings.Server.Prefix)).Subrouter()
	db := database.GetConnection()

	userRepository := repositories.UserRepository{DB: db}

	authHandler := handlers.AuthHandler{UserRepository: userRepository}

	sr.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	sr.HandleFunc("/auth/login", authHandler.Login).Methods("POST")

	return r
}
