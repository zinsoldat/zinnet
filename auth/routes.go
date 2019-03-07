package auth

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/zinsoldat/zinnet-go/auth/oauth"
	"github.com/zinsoldat/zinnet-go/models"
)

var oauthString = "random-string"

func index(e echo.Context) error {
	e.String(http.StatusOK, "/auth")
	return nil
}

// GetRoutes for auth handling
func GetRoutes() []models.Route {
	oauthProviders := []*oauth.OAuthProvider{
		oauth.GetGoogleProvider(),
		oauth.GetGithubProvider(),
	}

	routes := []models.Route{
		{Method: "GET", Path: "/auth", Handler: index},
	}
	for _, provider := range oauthProviders {
		routes = append(routes,
			models.Route{Method: "GET", Path: fmt.Sprintf("/auth/%s", provider.Name), Handler: provider.Redirect},
		)
		routes = append(routes,
			models.Route{Method: "GET", Path: fmt.Sprintf("/auth/%s/callback", provider.Name), Handler: provider.Callback},
		)
	}
	return routes
}
