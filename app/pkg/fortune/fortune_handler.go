package fortune

import (
	"go-practice/pkg/handlers"

	"net/http"

	"github.com/gin-gonic/gin"
)

type FortuneHandler struct{}

func NewFortuneHandler() *FortuneHandler {
	return &FortuneHandler{}
}

func (h *FortuneHandler) Routes() []handlers.Route {
	return []handlers.Route{
		{
			Method:  "GET",
			Path:    "/",
			Handler: h.Fortune,
		},
	}
}

func (h *FortuneHandler) Fortune(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": Fortune()})
}
