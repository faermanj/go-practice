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

// Register registers all handlers in the handlers package
func Register(r *gin.Engine) {
	// Add all handlers here
	handlers := []RouteHandler{
		NewHelloHandler(),
		NewMutantHandler(),
	}

	// Register all routes from all handlers
	for _, h := range handlers {
		for _, route := range h.Routes() {
			switch route.Method {
			case "GET":
				r.GET(route.Path, route.Handler)
			case "POST":
				r.POST(route.Path, route.Handler)
			case "PUT":
				r.PUT(route.Path, route.Handler)
			case "DELETE":
				r.DELETE(route.Path, route.Handler)
			}
		}
	}
}
