package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mateusrlopez/go-market/clients"
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

	stripeClient := clients.GetStripeClient()

	orderRepository := repositories.OrderRepository{Collection: db.Collection("orders")}
	paymentRepository := repositories.PaymentRepository{Collection: db.Collection("payments"), StripeClient: stripeClient}
	productRepository := repositories.ProductRepository{Collection: db.Collection("products")}
	reviewRepository := repositories.ReviewRepository{Collection: db.Collection("reviews")}
	skuRepository := repositories.SkuRepository{Collection: db.Collection("skus")}
	tokenRepository := repositories.TokenRepository{DB: rdb}
	userRepository := repositories.UserRepository{Collection: db.Collection("users")}

	authHandler := handlers.AuthHandler{TokenRepository: tokenRepository, UserRepository: userRepository}
	orderHandler := handlers.OrderHandler{OrderRepository: orderRepository}
	paymentHandler := handlers.PaymentHandler{PaymentRepository: paymentRepository}
	productHandler := handlers.ProductHandler{ProductRepository: productRepository}
	reviewHandler := handlers.ReviewHandler{ReviewRepository: reviewRepository}
	skuHandler := handlers.SkuHandler{SkuRepository: skuRepository}

	authMiddleware := middlewares.AuthorizationMiddleware{TokenRepository: tokenRepository, UserRepository: userRepository}

	sr.HandleFunc("/auth/register", authHandler.Register).Methods(http.MethodPost)
	sr.HandleFunc("/auth/login", authHandler.Login).Methods(http.MethodPost)
	sr.HandleFunc("/auth/refresh", authMiddleware.RefreshMiddleware(authHandler.Refresh)).Methods(http.MethodPost)
	sr.HandleFunc("/auth/logout", authMiddleware.AccessMiddleware(authHandler.Logout)).Methods(http.MethodPost)
	sr.HandleFunc("/auth/me", authMiddleware.AccessMiddleware(authHandler.Me)).Methods(http.MethodGet)

	sr.HandleFunc("/orders", authMiddleware.AccessMiddleware(orderHandler.Index)).Methods(http.MethodGet)
	sr.HandleFunc("/orders", authMiddleware.AccessMiddleware(orderHandler.Create)).Methods(http.MethodPost)
	sr.HandleFunc("/orders/{id}", authMiddleware.AccessMiddleware(orderHandler.Get)).Methods(http.MethodGet)
	sr.HandleFunc("/orders/{id}", authMiddleware.AccessMiddleware(orderHandler.Update)).Methods(http.MethodPut, http.MethodPatch)
	sr.HandleFunc("/orders/{id}", authMiddleware.AccessMiddleware(orderHandler.Delete)).Methods(http.MethodDelete)

	sr.HandleFunc("/payments", authMiddleware.AccessMiddleware(paymentHandler.Index)).Methods(http.MethodGet)
	sr.HandleFunc("/payments", authMiddleware.AccessMiddleware(paymentHandler.Create)).Methods(http.MethodPost)
	sr.HandleFunc("/payments/{id}", authMiddleware.AccessMiddleware(paymentHandler.Get)).Methods(http.MethodGet)

	sr.HandleFunc("/products", authMiddleware.AccessMiddleware(productHandler.Index)).Methods(http.MethodGet)
	sr.HandleFunc("/products", authMiddleware.AccessMiddleware(middlewares.AdminMiddleware(productHandler.Create))).Methods(http.MethodPost)
	sr.HandleFunc("/products/{id}", authMiddleware.AccessMiddleware(productHandler.Get)).Methods(http.MethodGet)
	sr.HandleFunc("/products/{id}", authMiddleware.AccessMiddleware(middlewares.AdminMiddleware(productHandler.Update))).Methods(http.MethodPut, http.MethodPatch)
	sr.HandleFunc("/products/{id}", authMiddleware.AccessMiddleware(middlewares.AdminMiddleware(productHandler.Delete))).Methods(http.MethodDelete)

	sr.HandleFunc("/reviews", authMiddleware.AccessMiddleware(reviewHandler.Index)).Methods(http.MethodGet)
	sr.HandleFunc("/reviews", authMiddleware.AccessMiddleware(reviewHandler.Create)).Methods(http.MethodPost)
	sr.HandleFunc("/reviews/{id}", authMiddleware.AccessMiddleware(reviewHandler.Update)).Methods(http.MethodPut, http.MethodPatch)
	sr.HandleFunc("/reviews/{id}", authMiddleware.AccessMiddleware(reviewHandler.Delete)).Methods(http.MethodDelete)

	sr.HandleFunc("/skus", authMiddleware.AccessMiddleware(skuHandler.Index)).Methods(http.MethodGet)
	sr.HandleFunc("/skus", authMiddleware.AccessMiddleware(middlewares.AdminMiddleware(skuHandler.Create))).Methods(http.MethodPost)
	sr.HandleFunc("/skus/{id}", authMiddleware.AccessMiddleware(middlewares.AdminMiddleware(skuHandler.Update))).Methods(http.MethodPut, http.MethodPatch)
	sr.HandleFunc("/skus/{id}", authMiddleware.AccessMiddleware(middlewares.AdminMiddleware(skuHandler.Delete))).Methods(http.MethodDelete)

	return r
}
