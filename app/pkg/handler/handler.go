package handlers

import (
	hs "go-practice/pkg/handlers"

	"go-practice/pkg/fortune"

	"github.com/gin-gonic/gin"

	"go-practice/pkg/hello"
	"go-practice/pkg/mutant"
)

// Register registers all handlers in the handlers package
func Register(r *gin.Engine) {
	// Add all handlers here
	handlers := []hs.RouteHandler{
		fortune.NewFortuneHandler(),
		hello.NewHelloHandler(),
		mutant.NewMutantHandler(),
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
