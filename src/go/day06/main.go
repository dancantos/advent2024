package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/dancantos/advent2024/src/go/pkg/timeit"
)

//go:embed input.txt
var input []byte
var maze, start = readInput(bytes.NewReader(input))

func readInput(r io.Reader) (sparseGrid, vec) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	y := 0
	grid := newSparseGrid(vec{})
	var start vec
	for s.Scan() {
		if y == 0 {
			grid.size.x = len(s.Text())
		}
		for x, char := range s.Text() {
			switch char {
			case '.':
			case '^':
				start = vec{x, y}
			case '#':
				grid.obstacles[vec{x, y}] = struct{}{}
			}
		}
		y++
	}
	grid.size.y = y
	return grid, start
}

func main() {
	const testInput = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	// maze, start := readInput(strings.NewReader(testInput))
	// maze, start := readInput(bytes.NewReader(input))

	fmt.Printf("Puzzle 1: count visited: %d\n", countVisited(maze, state{start, vec{0, -1}}))
	for range timeit.Run(1) {
		countVisited(maze, state{start, vec{0, -1}})
	}

	fmt.Printf("Puzzle 2: count viable blockades: %d\n", countViableBlockades(maze, state{start, vec{0, -1}}))
	for range timeit.Run(1) {
		countViableBlockades(maze, state{start, vec{0, -1}})
	}
}

func countVisited(grid sparseGrid, start state) int {
	path := newPath(grid.size)
	path.visit(start.p)
	grid.walk(start, func(s state) {
		path.visit(s.p)
	})
	return path.countVisited()
}

func countViableBlockades(grid sparseGrid, start state) int {
	path := newPath(grid.size)
	path.visit(start.p)
	count := 0
	grid.walk(start, func(s state) {
		if path.hasVisited(s.p) {
			return
		}
		path.visit(s.p)
		grid.obstacles[s.p] = struct{}{}
		if !terminates(grid, backtrack(s)) {
			count++
		}
		delete(grid.obstacles, s.p)
	})
	return count
}

func terminates(grid sparseGrid, start state) bool {
	return grid.walk(start, func(s state) {})
}
