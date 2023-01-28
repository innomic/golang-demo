package ch02

import (
	"mathx"
)

func LengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	for i, num := range nums {
		for j := 0; j < i; j++ {
			if num > nums[j] {
				dp[i] = mathx.Max(dp[i], dp[j]+1)
			}
		}
	}

	return mathx.Max(0, dp...)
}
