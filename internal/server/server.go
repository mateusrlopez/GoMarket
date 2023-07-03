package server

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mateusrlopez/go-market/internal/clients"
	"github.com/mateusrlopez/go-market/internal/configurations"
	"github.com/mateusrlopez/go-market/internal/controllers"
	"github.com/mateusrlopez/go-market/internal/middlewares"
	"github.com/mateusrlopez/go-market/internal/repositories"
	"github.com/mateusrlopez/go-market/internal/services"
)

func New() *http.Server {
	router := gin.Default()

	router.Use(cors.Default())

	jwtConfiguration := configurations.NewJwtConfiguration()
	mongoConfiguration := configurations.NewMongoConfiguration()
	redisConfiguration := configurations.NewRedisConfiguration()
	serverConfiguration := configurations.NewServerConfiguration()

	mongoClient := clients.NewMongoClient(mongoConfiguration)
	redisClient := clients.NewRedisClient(redisConfiguration)

	usersRepository := repositories.NewUsersRepository(mongoClient)
	tokenRepository := repositories.NewTokenRepository(redisClient)
	productsRepository := repositories.NewProductsRepository(mongoClient)
	reviewsRepository := repositories.NewReviewsRepository(mongoClient)
	ordersRepository := repositories.NewOrdersRepository(mongoClient)

	usersService := services.NewUsersService(usersRepository)
	tokenService := services.NewTokenService(jwtConfiguration, tokenRepository)
	authService := services.NewAuthService(tokenService, usersService)
	productsService := services.NewProductsService(productsRepository)
	reviewsService := services.NewReviewsService(reviewsRepository)
	ordersService := services.NewOrdersService(ordersRepository)

	authController := controllers.NewAuthController(authService)
	usersController := controllers.NewUsersController(usersService)
	productsController := controllers.NewProductsController(productsService)
	reviewsController := controllers.NewReviewsController(reviewsService)
	ordersController := controllers.NewOrdersController(ordersService)

	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
			auth.GET("/me", middlewares.AuthenticationMiddleware(authService), authController.Me)
			auth.DELETE("/logout", middlewares.AuthenticationMiddleware(authService), authController.Logout)
		}

		users := v1.Group("/users").Use(middlewares.AuthenticationMiddleware(authService))
		{
			users.GET("/", usersController.Index)
			users.GET("/:id", usersController.GetOneById)
			users.PUT("/:id", usersController.UpdateOneByID)
			users.DELETE("/:id", usersController.RemoveOneByID)
		}

		products := v1.Group("/products").Use(middlewares.AuthenticationMiddleware(authService))
		{
			products.POST("/", productsController.Create)
			products.GET("/", productsController.Index)
			products.GET("/:id", productsController.GetOneById)
			products.PUT("/:id", productsController.UpdateOneById)
			products.DELETE("/:id", productsController.RemoveOneById)
		}

		reviews := v1.Group("/reviews").Use(middlewares.AuthenticationMiddleware(authService))
		{
			reviews.POST("/", reviewsController.Create)
			reviews.GET("/", reviewsController.Index)
			reviews.GET("/:id", reviewsController.GetOneById)
			reviews.PUT("/:id", reviewsController.UpdateOneByID)
			reviews.DELETE("/:id", reviewsController.RemoveOneByID)
		}

		orders := v1.Group("/orders").Use(middlewares.AuthenticationMiddleware(authService))
		{
			orders.POST("/", ordersController.Create)
			orders.GET("/", ordersController.Index)
			orders.GET("/:id", ordersController.GetOneByID)
			orders.PUT("/:id", ordersController.UpdateOneByID)
			orders.DELETE("/:id", ordersController.RemoveOneByID)
		}
	}

	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%d", serverConfiguration.Port),
	}

	return server
}