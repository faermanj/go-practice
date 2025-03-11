package twoSum

import "testing"

func TestTwoSumMemo(t *testing.T) {
	values := []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 10
	expected := []int{11, -1}

	result := TwoSumMemo(values, target)
	if result == nil {
		t.Errorf("Expected result %v, but got nil", expected)
		return
	}
	matchOrdered := result[0] == expected[0] && result[1] == expected[1]
	matchUnordered := result[0] == expected[1] && result[1] == expected[0]
	var matches = matchOrdered || matchUnordered
	if (!matches) {
		t.Errorf("Expected result %v, but got %v", expected, result)
	}
}
