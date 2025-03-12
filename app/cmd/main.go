package main

import (
	"go-practice/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	handlers.Register(r)
	r.Run(":8080")
}
