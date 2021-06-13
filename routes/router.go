package routes

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/mateusrlopez/go-market/settings"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.PathPrefix(fmt.Sprintf("/%s/", settings.Settings.Server.Prefix))

	return r
}
