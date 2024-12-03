package main

import (
	_ "embed"
	"fmt"

	"github.com/dancantos/advent2024/src/go/pkg/timeit"
)

func main() {
	fmt.Printf("Puzzle1 (mulall): %d\n", mulall(input))
	for range timeit.Run(100) {
		mulall(input)
	}

	fmt.Println()

	fmt.Printf("Puzzle2 (mulWithInstruction): %d\n", mulWithInstruction(input))
	for range timeit.Run(100) {
		mulWithInstruction(input)
	}
}
