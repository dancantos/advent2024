package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"iter"

	_ "embed"
)

//go:embed input.txt
var input []byte
var g = readInput(bytes.NewReader(input))

func readInput(r io.Reader) grid[rune] {
	result := make(grid[rune], 0)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		if len(s.Text()) == 0 {
			continue
		}
		result = append(result, []rune(s.Text()))
	}
	return result
}

func main() {
	const input = `
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
	// g := readInput(strings.NewReader(input))
	// fmt.Println(findAllGeometry(g))
	fmt.Println(findAllGeometry2(g))
}

func findAllGeometry(g grid[rune]) int {
	searcher := gridSearcher{
		visited: map[vec]struct{}{},
	}
	result := 0
	for i := 0; i < g.width(); i++ {
		for j := 0; j < g.height(); j++ {
			start := vec{i, j}
			if _, visited := searcher.visited[vec{i, j}]; visited {
				continue
			}
			searcher.horizon = []vec{start}
			area, perimeter := bfs(g, &searcher)
			// fmt.Println(start, string(g.get(start)), area, perimeter)
			result += area * perimeter
		}
	}
	return result
}

func findAllGeometry2(g grid[rune]) int {
	searcher := gridSearcher{
		visited: map[vec]struct{}{},
	}
	result := 0
	for i := 0; i < g.width(); i++ {
		for j := 0; j < g.height(); j++ {
			start := vec{i, j}
			if _, visited := searcher.visited[start]; visited {
				continue
			}
			searcher.horizon = []vec{start}
			area, edges := bfs2(g, &searcher)
			fmt.Println(start, string(g.get(start)), area, edges)
			result += area * edges
		}
	}
	return result
}

func recordGeometry(area, perimeter *int) func(g grid[rune], current vec, visited map[vec]struct{}) iter.Seq[vec] {
	return func(g grid[rune], current vec, visited map[vec]struct{}) iter.Seq[vec] {
		return func(yield func(vec) bool) {
			*area++
			char := g.get(current)
			if current.x > 0 {
				left := vec{current.x - 1, current.y}
				if g.get(left) != char {
					*perimeter++
				} else if _, visited := visited[left]; !visited {
					yield(left)
				}
			} else {
				*perimeter++ // left boundary
			}

			if current.x < g.width()-1 {
				right := vec{current.x + 1, current.y}
				if g.get(right) != char {
					*perimeter++
				} else if _, visited := visited[right]; !visited {
					yield(right)
				}
			} else {
				*perimeter++ // right boundary
			}

			if current.y > 0 {
				top := vec{current.x, current.y - 1}
				if g.get(top) != char {
					*perimeter++
				} else if _, visited := visited[top]; !visited {
					yield(top)
				}
			} else {
				*perimeter++ // top boundary
			}

			if current.y < g.height()-1 {
				bottom := vec{current.x, current.y + 1}
				if g.get(bottom) != char {
					*perimeter++
				} else if _, visited := visited[bottom]; !visited {
					yield(bottom)
				}
			} else {
				*perimeter++ // bottom boundary
			}
		}
	}
}

func recordGeometry2(area *int, cornerGrid grid[int]) func(g grid[rune], current vec, visited map[vec]struct{}) iter.Seq[vec] {
	return func(g grid[rune], current vec, visited map[vec]struct{}) iter.Seq[vec] {
		return func(yield func(vec) bool) {
			*area++
			char := g.get(current)
			cornerGrid.set(current, cornerGrid.get(current)+1)
			cornerGrid.set(vec{current.x + 1, current.y}, cornerGrid.get(vec{current.x + 1, current.y})+1)
			cornerGrid.set(vec{current.x, current.y + 1}, cornerGrid.get(vec{current.x, current.y + 1})+1)
			cornerGrid.set(vec{current.x + 1, current.y + 1}, cornerGrid.get(vec{current.x + 1, current.y + 1})+1)
			l, r, t, d := vec{current.x - 1, current.y}, vec{current.x + 1, current.y}, vec{current.x, current.y - 1}, vec{current.x, current.y + 1}
			if g.inside(l) && g.get(l) == char {
				if _, visited := visited[l]; !visited {
					yield(l)
				}
			}
			if g.inside(r) && g.get(r) == char {
				if _, visited := visited[r]; !visited {
					yield(r)
				}
			}
			if g.inside(t) && g.get(t) == char {
				if _, visited := visited[t]; !visited {
					yield(t)
				}
			}
			if g.inside(d) && g.get(d) == char {
				if _, visited := visited[d]; !visited {
					yield(d)
				}
			}
		}
	}
}

func bfs(g grid[rune], searcher *gridSearcher) (int, int) {
	var area, perimeter int
	for len(searcher.horizon) != 0 {
		searcher.visit(g, searcher.pop(), recordGeometry(&area, &perimeter))
	}
	return area, perimeter
}

func bfs2(g grid[rune], searcher *gridSearcher) (int, int) {
	var area int
	cornerGrid := newGrid[int](vec{g.width() + 1, g.height() + 1})
	for len(searcher.horizon) != 0 {
		searcher.visit(g, searcher.pop(), recordGeometry2(&area, cornerGrid))
	}
	var edges int
	// var currentY int
	// for c, entry := range cornerGrid.all() {
	// 	if c.y > currentY {
	// 		fmt.Println()
	// 		currentY = c.y
	// 	}
	// 	if entry == 2 {
	// 		if c.x > 0 && c.y > 0 && c.x < g.width() && c.y < g.height() &&
	// 			g.get(c) == g.get(vec{c.x - 1, c.y - 1}) {
	// 			// cross case, there are 2 corners here
	// 			edges += 2
	// 			fmt.Print("C")
	// 		} else {
	// 			fmt.Print("+")
	// 		}
	// 	} else if entry%2 == 1 {
	// 		fmt.Print("C")
	// 		edges++
	// 	} else {
	// 		fmt.Print(".")
	// 	}
	// }
	// fmt.Println()
	// return area, edges

	for c, entry := range cornerGrid.all() {
		if entry == 2 {
			if c.x > 0 && c.y > 0 && c.x < g.width() && c.y < g.height() &&
				g.get(c) == g.get(vec{c.x - 1, c.y - 1}) {
				// cross case, there are 2 corners here
				edges += 2
			}
		} else if entry%2 == 1 {
			edges++
		}
	}
	return area, edges
}

type gridSearcher struct {
	visited map[vec]struct{}
	horizon []vec
}

func (g *gridSearcher) pop() vec {
	popped := g.horizon[0]
	g.horizon = g.horizon[1:]
	return popped
}

func (g *gridSearcher) visit(grid grid[rune], p vec, nextIter func(g grid[rune], current vec, visited map[vec]struct{}) iter.Seq[vec]) {
	if _, visited := g.visited[p]; visited {
		return
	}
	g.visited[p] = struct{}{}
	for next := range nextIter(grid, p, g.visited) {
		g.horizon = append(g.horizon, next)
	}
}

type (
	grid[T any] [][]T
	vec         struct{ x, y int }
)

func newGrid[T any](size vec) grid[T] {
	result := make([][]T, 0, size.y)
	for i := 0; i < size.y; i++ {
		result = append(result, make([]T, size.x))
	}
	return result
}

func (g grid[T]) all() iter.Seq2[vec, T] {
	return func(yield func(vec, T) bool) {
		for y := 0; y < g.height(); y++ {
			for x := 0; x < g.width(); x++ {
				if !yield(vec{x, y}, g[y][x]) {
					return
				}
			}
		}
	}
}

func (g grid[T]) width() int {
	return len(g[0])
}

func (g grid[T]) height() int {
	return len(g)
}

func (g grid[T]) get(p vec) T {
	return g[p.y][p.x]
}

func (g grid[T]) set(p vec, t T) {
	g[p.y][p.x] = t
}

func (g grid[T]) inside(p vec) bool {
	return 0 <= p.y && p.y < len(g) && 0 <= p.x && p.x < len(g[0])
}
