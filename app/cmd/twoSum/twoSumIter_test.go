package twoSum

import "testing"

func TestTwoSumIter(t *testing.T) {
	values := []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 10
	expected := []int{11, -1}

	result := TwoSumIter(values, target)
	if result == nil {
		t.Errorf("Expected result %v, but got nil", expected)
		return
	}
	if result[0] != expected[0] || result[1] != expected[1] {
		t.Errorf("Expected result %v, but got %v", expected, result)
	}
}
