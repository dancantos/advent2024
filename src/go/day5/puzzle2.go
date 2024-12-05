package main

import (
	"slices"
)

// if an array is not sorted, sort it and return the middle element
func badCounter(rules map[int][]int) func(arr []int) int {
	sortFn := sort(rules)
	return func(arr []int) int {
		if !isSorted(arr, rules) {
			slices.SortFunc(arr, sortFn)
			return arr[len(arr)/2]
		}
		return 0
	}
}
