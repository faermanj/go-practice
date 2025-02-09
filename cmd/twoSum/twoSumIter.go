package main

func twoSumIter(values []int, target int) []int {
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i]+values[j] == target {
				return []int{values[i], values[j]}
			}
		}
	}
	return nil
}
