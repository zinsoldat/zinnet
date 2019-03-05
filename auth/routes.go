package auth

import (
	"net/http"

	"github.com/zinsoldat/zinnet-go/models"
)

var oauthString = "random-string"

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("auth"))
}

// GetRoutes for auth handling
func GetRoutes() []models.Route {
	return []models.Route{
		{Path: "/auth", Handler: index},
		{Path: "/auth/github", Handler: redirectGithub},
		{Path: "/auth/github/callback", Handler: callbackGithub},
		{Path: "/auth/google", Handler: redirectGoogle},
		{Path: "/auth/google/callback", Handler: callbackGoogle},
	}
}
