package main

import (
	_ "embed"
)

func main() {
	puzzle1()
	puzzle2()
}

const (
	decreasing = -1
	increasing = 1
)

func countSafe(input [][]int, safeFn func([]int) bool) int {
	count := 0
	for _, nums := range input {
		if safeFn(nums) {
			count++
		}
	}
	return count
}
