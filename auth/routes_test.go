package auth_test

import (
	"testing"

	"github.com/zinsoldat/zinnet-go/auth"
)

func TestGetRoutes(t *testing.T) {
	routes := auth.GetRoutes()
	if len(routes) < 4 {
		t.Errorf("routes should contain at least 4 routes (oauth endpoints). It contains %d", len(routes))
	}
}
