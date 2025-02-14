package mutant

import (
	"fmt"
)

func checkOne(dna []string, curr string, x int, y int) bool {
	if (x < 0 || y < 0 || x >= len(dna) || y >= len(dna[x])) {
		return false
	}
 	if ! (curr == "T" || curr == "A" || curr == "C" || curr == "G") {
        return false
    }
	return string(dna[x][y]) == curr
}

func checkDNA(dna []string, curr string, 
		x0 int, y0 int, 
		x1 int, y1 int,
		x2 int, y2 int) bool {

	allMatch := checkOne(dna, curr, x0, y0) && checkOne(dna, curr, x1, y1) && checkOne(dna, curr, x2, y2)
	return  allMatch

}
	
func matchDNA(dna []string, i int, j int, curr string) bool {
	right := checkDNA(dna, curr, i, j + 1,  i, j + 2,  i, j + 3)
	down  := checkDNA(dna, curr, i+1, j,  i+2, j,  i+3, j)
	downR := checkDNA(dna, curr, i+1, j+1,  i+2, j+2,  i+3, j+3)
	downL := checkDNA(dna, curr, i+1, j-1,  i+2, j-2,  i+3, j-3)
	anyMatch := right || down || downR || downL
	return anyMatch
}

func IsMutant(dna []string) bool {
	for i, row := range dna {
		for j, col := range row {
			curr := string(col)
			fmt.Printf("i: %d, j: %d, col: %s\n", i, j, curr)
			if (matchDNA(dna, i, j, curr)) {
				fmt.Printf("Matched DNA: %s\n", curr)
				return true;
			}
		}
	}
	return false
}