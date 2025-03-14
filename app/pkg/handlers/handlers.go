package handlers

import (
	"github.com/gin-gonic/gin"
)

// RouteHandler interface defines what each handler must implement
type RouteHandler interface {
	Routes() []Route
}

// Route defines a single route with its method, path and handler
type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}
