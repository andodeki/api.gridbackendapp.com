package v1

import (
	"log"
	"net/http"

	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/app/router/routes"
	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/http/v1/httpcars"
	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/http/v1/httpusers"
	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/http/v1/version"
	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/repository/db"
	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/services/carservice"
	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/services/userservice"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-App-Token")
		if len(token) < 1 {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		log.Println("Inside V1 Middleware")

		next.ServeHTTP(w, r)
	})
}

func GetRoutes() (SubRoute map[string]routes.SubRoutePackage) {
	var atHandlerUsers = httpusers.NewUserHandler(userservice.NewUserService(db.NewUserRepository()))
	var atHandlerCars = httpcars.NewCarHandler(carservice.NewCarService(db.NewCarRepository()))

	/* ROUTES */
	SubRoute = map[string]routes.SubRoutePackage{
		"/v1": {
			Routes: routes.Routes{
				routes.Route{"Version", "GET", "/version", version.VersionHttp.VersionHandler},
				routes.Route{"UserCreate", "POST", "/user", atHandlerUsers.Create},
				routes.Route{"LoginUser", "POST", "/login", atHandlerUsers.Login},
				
				routes.Route{"CarCreate", "POST", "/car", atHandlerCars.CreateCar},
				// routes.Route{"UsersIndex", "GET", "/users", UsersHandler.Index},
				// routes.Route{"UsersStore", "POST", "/users", UsersHandler.Store},
				// routes.Route{"UsersEdit", "GET", "/users/{id}/edit", UsersHandler.Edit},
				// routes.Route{"UsersUpdate", "PUT", "/users/{id}", UsersHandler.Update},
				// routes.Route{"UsersDestroy", "DELETE", "/users/{id}", UsersHandler.Destroy},
			},
			// Middleware: Middleware,
		},
	}

	return
}
