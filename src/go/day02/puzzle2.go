package main

import (
	"fmt"
	"slices"
)

func puzzle2() {
	fmt.Println(countSafe(Input, isSafe2))
}

func isSafe2(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	if nums[0] < nums[len(nums)-1] {
		return _isSafe2(nums, increasing, false)
	}
	return _isSafe2(nums, decreasing, false)
}

func _isSafe2(nums []int, direction int, dampened bool) bool {
	for i := 0; i < len(nums)-1; i++ {
		diff := direction * (nums[i+1] - nums[i])
		if 1 <= diff && diff <= 3 {
			continue
		}
		if dampened {
			return false
		}
		return _isSafe2(slices.Concat(nums[:i], nums[i+1:]), direction, true) ||
			_isSafe2(slices.Concat(nums[:i+1], nums[i+2:]), direction, true)
	}
	return true
}
