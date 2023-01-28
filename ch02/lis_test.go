package ch02_test

import (
	"ch02"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLIS(t *testing.T) {
	for i, testcase := range []struct {
		nums []int
		res  int
	}{
		{
			nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			res:  4,
		},
		{
			nums: []int{1, 4, 3, 4, 2},
			res:  3,
		},
	} {
		res := ch02.LengthOfLIS(testcase.nums)
		assert.Equal(t, testcase.res, res, i)
	}
}
