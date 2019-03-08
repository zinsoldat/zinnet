package main

import (
	"context"
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/zinsoldat/zinnet-go/auth"
	"github.com/zinsoldat/zinnet-go/models"
	"github.com/zinsoldat/zinnet-go/users"
)

// Server struct
type Server struct {
	e       *echo.Echo
	config  *ServerConfig
	context context.Context
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
	setRoutes(e)

	return Server{
		e:       e,
		config:  config,
		context: context.Background(),
	}
}

// Start the server with the given configuration
func (s *Server) Start() {
	log.Fatal(s.e.Start(fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)))
}

// Stop the running server gracefully and clean up
// - close open database connections
// - terminate open requests
func (s *Server) Stop() {
	fmt.Println("stopping server")
	s.e.Shutdown(s.context)
	fmt.Println("stopped server")
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
