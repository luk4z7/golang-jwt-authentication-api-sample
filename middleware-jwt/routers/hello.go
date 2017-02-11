package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"middleware-jwt/controllers"
	"middleware-jwt/core/authentication"
)

func SetHelloRoutes(router *mux.Router) *mux.Router {
	router.Handle("/test/hello", negroni.New(
		// handlers
		negroni.HandlerFunc(authentication.RequireTokenAuthentication),
		// middleware
		negroni.HandlerFunc(controllers.HelloController),
	)).Methods("GET")

	return router
}
