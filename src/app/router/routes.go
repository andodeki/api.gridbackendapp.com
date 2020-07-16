package router

import (
	"log"
	"net/http"

	ServerPing "github.com/andodeki/code/HA/api.gridbackendapp.com/src/http/v1/serverping"

	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/app/router/routes"
	// "github.com/andodeki/code/HA/api.griffins.com/src/app/router/routes"
	// ServerPing "github.com/andodeki/code/HA/api.griffins.com/src/http/v1/serverping"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Inside main middleware.")
		next.ServeHTTP(w, r)
	})
}

func GetRoutes() routes.Routes {

	return routes.Routes{
		routes.Route{"Ping", "GET", "/ping", ServerPing.PingHttp.Ping},
	}
}
