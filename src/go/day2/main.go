package main

import (
	_ "embed"
	"slices"

	"github.com/dancantos/advent2025/src/go/pkg/it"
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
	return it.Count(it.Filter(slices.Values(input), safeFn))
}
