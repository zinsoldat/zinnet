package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/zinsoldat/zinnet-go/auth"
	"github.com/zinsoldat/zinnet-go/models"
	"github.com/zinsoldat/zinnet-go/users"
)

// Server struct
type Server struct {
	server *http.Server
}

// ServerConfig struct to configure a server
type ServerConfig struct {
	Port int16
	Host string
}

// Start the server with the given configuration
func (s *Server) Start(config ServerConfig) {
	router := mux.NewRouter()
	initlizeRoutes(router)

	s.server = &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
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
