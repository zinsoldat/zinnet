package users

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/zinsoldat/zinnet-go/models"
)

func getAll(e echo.Context) error {
	e.String(http.StatusOK, "get all users")
	return nil
}

// GetRoutes for users handling
func GetRoutes() []models.Route {
	return []models.Route{
		{Path: "/users", Handler: getAll},
	}
}
