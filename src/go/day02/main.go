package main

import (
	_ "embed"
	"fmt"
	"slices"

	"github.com/dancantos/advent2024/src/go/pkg/it"
	"github.com/dancantos/advent2024/src/go/pkg/timeit"
)

func main() {
	fmt.Printf("Puzzle1 (isSafe): %d\n", countSafe(Input, isSafe1))
	for range timeit.Run(100) {
		countSafe(Input, isSafe1)
	}

	fmt.Println()

	fmt.Printf("Puzzle2 (isSafeDampened): %d\n", countSafe(Input, isSafe2))
	for range timeit.Run(100) {
		countSafe(Input, isSafe2)
	}
}

const (
	decreasing = -1
	increasing = 1
)

func countSafe(input [][]int, safeFn func([]int) bool) int {
	return it.Count(it.Filter(slices.Values(input), safeFn))
}
