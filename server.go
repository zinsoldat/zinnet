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

// ServerConfig struct to configure a `Server` with the `NewServer` method
type ServerConfig struct {
	Port int16
	Host string
}

// NewServer creates a new server struct with the given `ServerConfig`.
// During creation the routes of the server are set.
func NewServer(config *ServerConfig) Server {

	router := mux.NewRouter()
	setRoutes(router)

	return Server{
		server: &http.Server{
			Handler:      router,
			Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port),
			WriteTimeout: 5 * time.Second,
			ReadTimeout:  5 * time.Second,
		},
	}
}

// Start the server with the given configuration
func (s *Server) Start() {
	fmt.Printf("server starting on %s\n", s.server.Addr)
	s.server.ListenAndServe()
}

// Stop the running server gracefully and clean up
// - close open database connections
// - terminate open requests
func (s *Server) Stop() {
	fmt.Printf("stopping server on %s\n", s.server.Addr)
	s.server.Close()
	fmt.Printf("stopped server on %s\n", s.server.Addr)
}

func setRoutes(r *mux.Router) {
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
