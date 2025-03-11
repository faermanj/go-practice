package main

import (
	"log"
    "github.com/gin-gonic/gin"

	"go-practice/pkg/hello"
	"go-practice/pkg/twoSum"
	"go-practice/pkg/mutant"
	"go-practice/internal/handlers"
)

func main() {
	r := gin.Default()

	// Initialize and register handlers
	h := handlers.NewHandler()
	h.Register(r)

	// Example code below
	isMutant := mutant.IsMutant([]string{
		"ATGCGA", 
		"CAGTGC", 
		"TTATGT", 
		"AGATTG", 
		"TCCCTA", 
		"TCACTG"})
	log.Printf("Is mutant? %t\n", isMutant)

	log.Printf("%s", hello.Hello())

	values := make([]int, 1_000_000_000)
	values[0] = 1
	values[999_999_999] = 1336
	target := 1337
	result := twoSum.TwoSumIter(values, target)
	log.Printf("Found %d = %d + %d\n", target, result[0], result[1])
	result = twoSum.TwoSumMemo(values, target)
	log.Printf("Found %d = %d + %d\n", target, result[0], result[1])
	result = twoSum.TwoSumWin(values, target)
	log.Printf("Found %d = %d + %d\n", target, result[0], result[1])
}
