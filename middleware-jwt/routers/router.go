package routers

import (
	"github.com/gorilla/mux"
)

// Initialize the routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetHelloRoutes(router)
	router = SetAuthenticationRoutes(router)
	return router
}
