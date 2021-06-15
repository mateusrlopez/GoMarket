package routes

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/mateusrlopez/go-market/database"
	"github.com/mateusrlopez/go-market/middlewares"
	"github.com/mateusrlopez/go-market/settings"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.Use(middlewares.HeadersMiddleware)

	sr := r.PathPrefix(fmt.Sprintf("/%s", settings.Settings.Server.Prefix)).Subrouter()
	db := database.GetConnection()

	setupAuthRoutes(sr, db)

	return r
}
