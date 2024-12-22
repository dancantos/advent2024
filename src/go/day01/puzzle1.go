package main

import (
	"fmt"
	"slices"
)

func puzzle1() {
	fmt.Println(sumdiffs(a1, a2))
}

func sumdiffs(a1, a2 []int) int {
	// a1 = slices.Clone(a1)
	// a2 = slices.Clone(a2)
	slices.Sort(a1)
	slices.Sort(a2)

	sum := 0
	for i := 0; i < len(a1); i++ {
		sum += abs(a1[i] - a2[i])
	}
	return sum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
