package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/mateusrlopez/go-market/routes"
	"github.com/mateusrlopez/go-market/settings"
	log "github.com/sirupsen/logrus"
	"github.com/vrischmann/envconfig"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading variables from .env file: %s", err)
	}

	if err := envconfig.Init(&settings.Settings); err != nil {
		log.Fatalf("Error initializing configuration: %s", err)
	}
}

func main() {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", settings.Settings.Server.Port),
		Handler: routes.SetupRoutes(),
	}

	log.Printf("Starting server on port %s", settings.Settings.Server.Port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
