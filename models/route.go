package models

import "github.com/labstack/echo"

type Route struct {
	Path    string
	Method  string
	Handler func(e echo.Context) error
}
