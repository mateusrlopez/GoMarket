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
	sr := r.PathPrefix(fmt.Sprintf("/%s", settings.Settings.Server.Prefix)).Subrouter()

	r.Use(middlewares.HeadersMiddleware)

	db := database.GetMongoConnection()
	rdb := database.GetRedisConnection()

	tokenRepository := repositories.TokenRepository{DB: rdb}
	userRepository := repositories.UserRepository{Collection: db.Collection("users")}

	authHandler := handlers.AuthHandler{TokenRepository: tokenRepository, UserRepository: userRepository}

	authMiddleware := middlewares.AuthorizationMiddleware{TokenRepository: tokenRepository, UserRepository: userRepository}

	sr.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	sr.HandleFunc("/auth/login", authHandler.Login).Methods("POST")
	sr.HandleFunc("/auth/logout", authMiddleware.AccessMiddleware(authHandler.Logout)).Methods("POST")
	sr.HandleFunc("/auth/me", authMiddleware.AccessMiddleware(authHandler.Me)).Methods("GET")

	return r
}
