package users

import (
	"net/http"

	"github.com/zinsoldat/zinnet-go/models"
)

func getAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get all users"))
}

// GetRoutes for users handling
func GetRoutes() []models.Route {
	return []models.Route{
		{Path: "/users", Handler: getAll},
	}
}
