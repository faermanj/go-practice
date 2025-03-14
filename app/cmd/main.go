package main

import (
	h "go-practice/pkg/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	h.Register(r)
	r.Run(":8080")
}
