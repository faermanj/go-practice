package twoSum_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTwoSum(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Two Sum - Iterative - Test Suite")
}
