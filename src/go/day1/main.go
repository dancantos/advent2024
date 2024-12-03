package main

import (
	"fmt"
	"slices"

	"github.com/dancantos/advent2024/src/go/pkg/timeit"
)

func main() {
	c1, c2 := slices.Clone(a1), slices.Clone(a2)

	fmt.Printf("Puzzle1 (sumdiffs): %d\n", sumdiffs(c1, c2))
	for range timeit.Run(100) {
		sumdiffs(c1, c2)
	}

	fmt.Println()

	c1, c2 = slices.Clone(a1), slices.Clone(a2)
	fmt.Printf("Puzzle2 (similarityScore): %d\n", similarityScore(c1, c2))
	for range timeit.Run(100) {
		similarityScore(c1, c2)
	}
}
