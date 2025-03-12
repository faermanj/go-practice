package handlers

import (
	"go-practice/pkg/mutant"
	"net/http"

	"github.com/gin-gonic/gin"
)

// curl -vX POST http://localhost:8080/mutant -H "Content-Type: application/json"  -d '{ "dna": ["ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"]}'
// curl -vX POST http://localhost:8080/mutant -H "Content-Type: application/json"  -d '{ "dna": ["TTGCTA", "CAGTGC", "TTATGT", "AGAAGG", "CCCTTA", "TCACTG"]}'

type DNARequest struct {
	DNA []string `json:"dna"`
}

type MutantHandler struct{}

func NewMutantHandler() *MutantHandler {
	return &MutantHandler{}
}

func (h *MutantHandler) Routes() []Route {
	return []Route{
		{
			Method:  "POST",
			Path:    "/mutant",
			Handler: h.IsMutant,
		},
	}
}

func (h *MutantHandler) IsMutant(c *gin.Context) {
	var req DNARequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if mutant.IsMutant(req.DNA) {
		c.JSON(http.StatusOK, gin.H{"message": "Mutant DNA detected"})
	} else {
		c.JSON(http.StatusForbidden, gin.H{"message": "Not a mutant"})
	}
}
