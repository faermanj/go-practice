package twoSum

import (
	"sort"
)

func TwoSumWin(array []int, targetSum int) []int {
	if array == nil {
		return nil
	}
	sort.Ints(array) // O(n log n)
	left := 0
	right := len(array) - 1
	for left < right {
		sum := array[left] + array[right]
		if sum == targetSum {
			return []int{array[left], array[right]}
		} else if sum < targetSum {
			left++
		} else {
			right--
		}
	}
	return nil
}
