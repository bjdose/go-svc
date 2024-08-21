package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"testing"
)

func Test_routes_exist(t *testing.T) {
	testApp := Config{}

	testRoutes := testApp.routes()
	chiRoutes := testRoutes.(chi.Router)

	routes := []string{"/authenticate"}

	for _, route := range routes {
		routeExists(t, chiRoutes, route)
	}

}

func routeExists(t *testing.T, routes chi.Router, route string) {
	found := false

	// Iterate over the routes and check if the route exists
	_ = chi.Walk(routes, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if route == route {
			found = true
		}

		return nil
	})

	if !found {
		t.Errorf("route %s not found", route)
	}
}
