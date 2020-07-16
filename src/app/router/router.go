package router

import (
	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/app/router/routes"
	V1SubRoutes "github.com/andodeki/code/HA/api.gridbackendapp.com/src/http/v1"
	"github.com/gorilla/mux"
)

type Router struct {
	Router *mux.Router
}

func (r *Router) Init() {
	// r.Router.Use(Middleware)

	baseRoutes := GetRoutes()
	for _, route := range baseRoutes {
		r.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	v1SubRoutes := V1SubRoutes.GetRoutes()
	for name, pack := range v1SubRoutes {
		r.AttachSubRouterWithMiddleware(name, pack.Routes, pack.Middleware)
	}
}

func (r *Router) AttachSubRouterWithMiddleware(path string, subroutes routes.Routes, middleware mux.MiddlewareFunc) (SubRouter *mux.Router) {

	SubRouter = r.Router.PathPrefix(path).Subrouter()
	// SubRouter.Use(middleware)

	for _, route := range subroutes {
		SubRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return
}

func NewRouter() (r Router) {
	r.Router = mux.NewRouter().StrictSlash(true)

	return
}
