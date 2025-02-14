package main

import (
	"log"
    "github.com/gin-gonic/gin"

	"go-practice/cmd/hello"
	"go-practice/cmd/twoSum"
	"go-practice/cmd/mutant"
	"go-practice/cmd/handlers"

)

func main() {
	r := gin.Default()

    r.POST("/mutant", handlers.IsMutantHandler)

    r.Run(":8080") // Run on port 8080

	return;

	isMutant := mutant.IsMutant([]string{
		"ATGCGA", 
		"CAGTGC", 
		"TTATGT", 
		"AGATTG", 
		"TCCCTA", 
		"TCACTG"})
	log.Printf("Is mutant? %t\n", isMutant)


	// Hello, World!
	log.Printf("%s", hello.Hello())

	// Two Sum
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
	//
}
