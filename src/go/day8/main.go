package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"iter"

	"github.com/dancantos/advent2024/src/go/pkg/grid"
	"github.com/dancantos/advent2024/src/go/pkg/input"
	"github.com/dancantos/advent2024/src/go/pkg/it"
	"github.com/dancantos/advent2024/src/go/pkg/timeit"
)

//go:embed input.txt
var data []byte
var inputGrid, size = readInput(bytes.NewReader(data))

func readInput(r io.Reader) (grid2, vec) {
	result := make(grid2)
	width, height := input.ReadGrid(r, func(x, y int, char rune) {
		if char != '.' {
			result[char] = append(result[char], vec{x, y})
		}
	})
	return result, vec{width, height}
}

func main() {
	fmt.Println("Puzzle1 (Count 2 antinodes): ", countAntinodes(inputGrid, size, iterateAntinodes))
	for range timeit.Run(1) {
		countAntinodes(inputGrid, size, iterateAntinodes)
	}

	fmt.Println("Puzzle2 (Count all antinodes): ", countAntinodes(inputGrid, size, iterateMoreAntinodes))
	for range timeit.Run(1) {
		countAntinodes(inputGrid, size, iterateMoreAntinodes)
	}
}

func countAntinodes(g grid2, size vec, antinodes func(n1, n2, size vec) iter.Seq[vec]) int {
	visited := grid.NewBitmask(size.x, size.y)
	count := 0
	for _, nodes := range g {
		for n1, n2 := range it.SlicePairs(nodes) {
			for a := range antinodes(n1, n2, size) {
				if visited.Set(a.x, a.y) {
					count++
				}
			}
		}
	}
	return count
}

func iteratePairs[T any](items []T) iter.Seq2[T, T] {
	return func(yield func(T, T) bool) {
		for i := 0; i < len(items); i++ {
			for j := i + 1; j < len(items); j++ {
				if !yield(items[i], items[j]) {
					return
				}
			}
		}
	}
}

func iterateAntinodes(n1, n2 vec, size vec) iter.Seq[vec] {
	diff := vec{n2.x - n1.x, n2.y - n1.y}
	a1 := sub(n1, diff)
	a2 := add(n2, diff)
	return func(yield func(vec) bool) {
		if inside(a1, size) {
			if !yield(a1) {
				return
			}
		}
		if inside(a2, size) {
			if !yield(a2) {
				return
			}
		}
	}
}

func iterateMoreAntinodes(n1, n2 vec, size vec) iter.Seq[vec] {
	diff := vec{n2.x - n1.x, n2.y - n1.y}
	return func(yield func(vec) bool) {
		current := n1
		for inside(current, size) {
			if !yield(current) {
				return
			}
			current = sub(current, diff)
		}

		current = n2
		for inside(current, size) {
			if !yield(current) {
				return
			}
			current = add(current, diff)
		}
	}
}

func add(v1, v2 vec) vec {
	return vec{v1.x + v2.x, v1.y + v2.y}
}

func sub(v1, v2 vec) vec {
	return vec{v1.x - v2.x, v1.y - v2.y}
}

func inside(p, size vec) bool {
	return p.x >= 0 && p.x < size.x && p.y >= 0 && p.y < size.y
}

type (
	grid2 map[rune][]vec
	vec   struct{ x, y int }
)
