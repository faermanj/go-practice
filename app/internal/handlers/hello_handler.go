package handlers

import (
	"go-practice/pkg/hello"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloHandler struct{}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (h *HelloHandler) Routes() []Route {
	return []Route{
		{
			Method:  "GET",
			Path:    "/hello",
			Handler: h.HelloWorld,
		},
	}
}

func (h *HelloHandler) HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": hello.Hello()})
}
