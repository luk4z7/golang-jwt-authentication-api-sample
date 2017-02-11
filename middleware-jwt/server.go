package main

import (
	"github.com/codegangsta/negroni"
	"middleware-jwt/routers"
	"middleware-jwt/settings"
	"net/http"
)

func main() {
	settings.Init()
	// Get all routes defined
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":6060", n)
}
