package main

import (
	_ "embed"
	"fmt"

	"github.com/dancantos/advent2024/src/go/pkg/timeit"
)

func main() {
	var result1 int
	d1 := timeit.Run(func() {
		result1 = mulall(input)
	})

	var result2 int
	d2 := timeit.Run(func() {
		result2 = mulWithInstruction(input)
	})

	fmt.Printf("Puzzle1 (mulall): %d (in %s)\n", result1, d1)
	fmt.Printf("Puzzle2 (mulWithInstruction): %d (in %s)\n", result2, d2)
}
