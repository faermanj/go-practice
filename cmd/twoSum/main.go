package main

import (
	"fmt"
)

func main() {
	values := []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 10
	result := twoSumMemo(values, target)
	if result == nil {
		fmt.Println("No result found")
		return
	}
	fmt.Printf("Found %d = %d + %d\n", target, result[0], result[1])
}
