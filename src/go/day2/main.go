package main

import (
	_ "embed"
	"fmt"
	"slices"

	"github.com/dancantos/advent2024/src/go/pkg/it"
	"github.com/dancantos/advent2024/src/go/pkg/timeit"
)

func main() {
	var result1 int
	d1 := timeit.Run(func() {
		result1 = countSafe(Input, isSafe1)
	})

	var result2 int
	d2 := timeit.Run(func() {
		result2 = countSafe(Input, isSafe2)
	})

	fmt.Printf("Puzzle1: %d (in %s)\n", result1, d1)
	fmt.Printf("Puzzle2: %d (in %s)\n", result2, d2)
}

const (
	decreasing = -1
	increasing = 1
)

func countSafe(input [][]int, safeFn func([]int) bool) int {
	return it.Count(it.Filter(slices.Values(input), safeFn))
}
