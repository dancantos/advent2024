package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"iter"
)

//go:embed input.txt
var input []byte

var g, starts = readInput(bytes.NewReader(input))

func readInput(r io.Reader) (grid, []vec) {
	result := [][]int{}
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	width, y := 0, 0
	zeroes := []vec{}
	for s.Scan() {
		width = len(s.Text())
		row := make([]int, width)
		for x, char := range s.Text() {
			row[x] = int(char - '0')
			if char == '0' {
				zeroes = append(zeroes, vec{x, y})
			}
		}
		result = append(result, row)
		y++
	}
	return grid(result), zeroes
}

func main() {
	fmt.Println(puzzle1(g, starts))
}

func puzzle1(g grid, starts []vec) int {
	searcher := &graphSearcher{}
	count := 0
	for _, s := range starts {
		for node := range searcher.walk([]vec{s}, g, accept1) {
			if g.get(node) == 9 {
				count++
			}
		}
	}
	return count
}

type (
	grid [][]int
	vec  struct{ x, y int }
)

func (g grid) get(p vec) int {
	return g[p.y][p.x]
}

func (g grid) height() int {
	return len(g)
}

func (g grid) width() int {
	return len(g[0])
}

type graphSearcher struct {
	visited map[vec]struct{}
	queue   []vec
}

func (g *graphSearcher) pop() vec {
	result := g.queue[0]
	g.queue = g.queue[1:]
	return result
}

func (g *graphSearcher) walk(starts []vec, gr grid, accept func(grid, vec, vec) bool) iter.Seq[vec] {
	g.queue = starts
	g.visited = make(map[vec]struct{})
	return func(yield func(vec) bool) {
		for len(g.queue) > 0 {
			node := g.pop()
			if _, v := g.visited[node]; v {
				continue
			}
			if !yield(node) {
				return
			}
			g.visit(node, gr, accept)
		}
	}
}

func (g *graphSearcher) visit(node vec, grid grid, accept func(grid, vec, vec) bool) {
	g.visited[node] = struct{}{}
	if node.x > 0 {
		next := vec{node.x - 1, node.y}
		if accept(grid, node, next) {
			g.queue = append(g.queue, vec{node.x - 1, node.y})
		}
	}
	if node.x < grid.width()-1 {
		next := vec{node.x + 1, node.y}
		if accept(grid, node, next) {
			g.queue = append(g.queue, next)
		}
	}
	if node.y > 0 {
		next := vec{node.x, node.y - 1}
		if accept(grid, node, next) {
			g.queue = append(g.queue, next)
		}
	}
	if node.y < grid.height()-1 {
		next := vec{node.x, node.y + 1}
		if accept(grid, node, next) {
			g.queue = append(g.queue, next)
		}
	}
}

func accept1(g grid, current vec, next vec) bool {
	return g.get(current) == g.get(next)-1
}
