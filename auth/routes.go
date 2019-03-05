package auth

import (
	"net/http"

	"github.com/zinsoldat/zinnet-go/auth/oauth"
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
		{Path: "/auth/github", Handler: oauth.RedirectGithub},
		{Path: "/auth/github/callback", Handler: oauth.CallbackGithub},
		{Path: "/auth/google", Handler: oauth.RedirectGoogle},
		{Path: "/auth/google/callback", Handler: oauth.CallbackGoogle},
	}
}
