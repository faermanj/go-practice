package handlers

import (
	"github.com/gin-gonic/gin"
)

// Handler represents a collection of HTTP handlers with their dependencies
type Handler struct {
	// Add any dependencies here
	// db *sql.DB
	// services *Services
}

// NewHandler creates a new handler instance with dependencies
func NewHandler() *Handler {
	return &Handler{
		// Initialize dependencies
	}
}

// Register registers all routes for this handler
func (h *Handler) Register(r *gin.Engine) {
	// Register all routes here
	r.POST("/mutant", h.IsMutantHandler)
	// Add more routes as needed
}

// IsMutantHandler handles the mutant endpoint
func (h *Handler) IsMutantHandler(c *gin.Context) {
	// Your existing handler logic
} 