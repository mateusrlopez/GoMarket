package routes

import (
	"fmt"
	"net/http"

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
	productRepository := repositories.ProductRepository{Collection: db.Collection("products")}

	authHandler := handlers.AuthHandler{TokenRepository: tokenRepository, UserRepository: userRepository}
	productHandler := handlers.ProductHandler{ProductRepository: productRepository}

	authMiddleware := middlewares.AuthorizationMiddleware{TokenRepository: tokenRepository, UserRepository: userRepository}

	sr.HandleFunc("/auth/register", authHandler.Register).Methods(http.MethodPost)
	sr.HandleFunc("/auth/login", authHandler.Login).Methods(http.MethodPost)
	sr.HandleFunc("/auth/logout", authMiddleware.AccessMiddleware(authHandler.Logout)).Methods(http.MethodPost)
	sr.HandleFunc("/auth/me", authMiddleware.AccessMiddleware(authHandler.Me)).Methods(http.MethodGet)

	sr.HandleFunc("/products", authMiddleware.AccessMiddleware(productHandler.Index)).Methods(http.MethodGet)
	sr.HandleFunc("/products", authMiddleware.AccessMiddleware(middlewares.AdminMiddleware(productHandler.Create))).Methods(http.MethodPost)
	sr.HandleFunc("/products/{id}", authMiddleware.AccessMiddleware(productHandler.Get)).Methods(http.MethodGet)
	sr.HandleFunc("/products/{id}", authMiddleware.AccessMiddleware(middlewares.AdminMiddleware(productHandler.Update))).Methods(http.MethodPut)
	sr.HandleFunc("/products/{id}", authMiddleware.AccessMiddleware(middlewares.AdminMiddleware(productHandler.Delete))).Methods(http.MethodDelete)

	return r
}
