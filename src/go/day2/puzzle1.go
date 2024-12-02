package main

import "fmt"

func puzzle1() {
	fmt.Println(countSafe(Input, isSafe1))
}

func isSafe1(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	if nums[0] < nums[len(nums)-1] {
		return _isSafe1(nums, increasing)
	}
	return _isSafe1(nums, decreasing)
}

func _isSafe1(nums []int, direction int) bool {
	for i := 1; i < len(nums); i++ {
		diff := direction * (nums[i] - nums[i-1])
		if 1 <= diff && diff <= 3 {
			continue
		}
		return false
	}
	return true
}
