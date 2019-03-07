package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/labstack/echo"
	"github.com/zinsoldat/zinnet-go/auth"
	"github.com/zinsoldat/zinnet-go/models"
	"github.com/zinsoldat/zinnet-go/users"
)

// Server struct
type Server struct {
	server *http.Server
	e      *echo.Echo
	config *ServerConfig
}

// ServerConfig struct to configure a `Server` with the `NewServer` method
type ServerConfig struct {
	Port int16
	Host string
}

// NewServer creates a new server struct with the given `ServerConfig`.
// During creation the routes of the server are set.
func NewServer(config *ServerConfig) Server {
	e := echo.New()
	router := mux.NewRouter()
	setRoutes(e)

	return Server{
		e:      e,
		config: config,
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
	log.Fatal(s.e.Start(fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)))
}

// Stop the running server gracefully and clean up
// - close open database connections
// - terminate open requests
func (s *Server) Stop() {
	fmt.Printf("stopping server on %s\n", s.server.Addr)
	s.server.Close()
	fmt.Printf("stopped server on %s\n", s.server.Addr)
}

func setRoutes(r *echo.Echo) {
	routes := [][]models.Route{
		auth.GetRoutes(),
		users.GetRoutes(),
	}

	for _, subRoutings := range routes {
		for _, route := range subRoutings {
			r.Add(route.Method, route.Path, route.Handler)
		}
	}
}
