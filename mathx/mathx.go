package mathx

import (
	"golang.org/x/exp/constraints"
)

func Max[T constraints.Ordered](a T, nums ...T) T {
	if len(nums) == 0 {
		return a
	}
	head := nums[0]
	tail := nums[1:]

	b := Max(head, tail...)
	if a < b {
		return b
	}
	return a
}
