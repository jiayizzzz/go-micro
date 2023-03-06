package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"testing"
)

func Test_routes_exist(t *testing.T) {
	testApp := Config{}

	testRoutes := testApp.routes()
	chiRoutes := testRoutes.(chi.Routes)

	routes := []string{"/authenticate"}

	for _, routes := range routes {
		routeExists(t, chiRoutes, routes)
	}
}

func routeExists(t *testing.T, routes chi.Routes, route string) {
	found := false
	_ = chi.Walk(routes, func(method string, foundRoute string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if route == foundRoute {
			found = true
		}
		return nil
	})
	if !found {
		t.Errorf("didn't find %s in registered routes", route)
	}
}
