package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zinsoldat/zinnet-go/auth"
	"github.com/zinsoldat/zinnet-go/models"
	"github.com/zinsoldat/zinnet-go/users"
)

func main() {
	router := mux.NewRouter()
	initlizeRoutes(router)

	port := ":3000"
	fmt.Println("\nListening on port " + port)
	http.ListenAndServe(port, router)
}

func initlizeRoutes(r *mux.Router) {
	routes := [][]models.Route{
		auth.GetRoutes(),
		users.GetRoutes(),
	}

	for _, subRoutings := range routes {
		for _, route := range subRoutings {
			r.HandleFunc(route.Path, route.Handler)
		}
	}
}
