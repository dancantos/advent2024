package main

// if an array is sorted, return the middle element
func goodCounter(rules map[int][]int) func(arr []int) int {
	return func(arr []int) int {
		if isSorted(arr, rules) {
			return arr[len(arr)/2]
		}
		return 0
	}
}
