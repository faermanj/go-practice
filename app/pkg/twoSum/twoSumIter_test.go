package twoSum

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// unit test
func TestTwoSumIter(t *testing.T) {
	values := []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 10
	expected := []int{11, -1}

	result := TwoSumIter(values, target)
	if result == nil {
		t.Errorf("Expected result %v, but got nil", expected)
		return
	}
	matchOrdered := result[0] == expected[0] && result[1] == expected[1]
	matchUnordered := result[0] == expected[1] && result[1] == expected[0]
	var matches = matchOrdered || matchUnordered
	if !matches {
		t.Errorf("Expected result %v, but got %v", expected, result)
	}
}

// ginkgo test
var _ = Describe("TwoSumIter", func() {
	It("should return the correct values for the target sum", func() {
		values := []int{3, 5, -4, 8, 11, 1, -1, 6}
		target := 10
		expected := []int{11, -1}
		result := TwoSumIter(values, target)
		Expect(result).NotTo(BeNil(), "Expected result %v, but got nil", expected)
		Expect(result).To(ConsistOf(expected), "Expected result %v, but got %v", expected, result)
	})
})
