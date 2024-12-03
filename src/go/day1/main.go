package main

import (
	"fmt"
	"slices"

	"github.com/dancantos/advent2024/src/go/pkg/timeit"
)

func main() {
	a1, a2, err := readInput(data)
	if err != nil {
		panic(err)
	}
	c1, c2 := slices.Clone(a1), slices.Clone(a2)

	var result1 int
	d1 := timeit.Run(func() {
		result1 = sumdiffs(c1, c2)
	})

	c1, c2 = slices.Clone(a1), slices.Clone(a2)
	var result2 int
	d2 := timeit.Run(func() {
		result2 = similarityScore(c1, c2)
	})

	fmt.Printf("Puzzle1 (sumdiffs): %d (in %s)\n", result1, d1)
	fmt.Printf("Puzzle2 (similarityScore): %d (in %s)\n", result2, d2)
}
