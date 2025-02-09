package twoSum

func TwoSumMemo(values []int, target int) []int {
	memo := make(map[int]struct{})
	for _, value := range values {
		diff := target - value
		if _, ok := memo[diff]; ok {
			return []int{value, diff}
		}
		memo[value] = struct{}{}
	}
	return nil
}
