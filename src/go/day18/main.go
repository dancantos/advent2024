package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/dancantos/advent2024/src/go/pkg/must"
)

//go:embed input.txt
var inputTxt []byte
var input = readInput(bytes.NewReader(inputTxt))

func readInput(r io.Reader) []vec {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	result := []vec{}
	for s.Scan() {
		if s.Text() == "" {
			continue
		}
		s := strings.TrimSpace(s.Text())
		split := strings.Split(s, ",")
		result = append(result, vec{must.Return(strconv.Atoi(split[0])), must.Return(strconv.Atoi(split[1]))})
	}
	return result
}

func main() {
	size := vec{71, 71}
	end := vec{70, 70}
	start := vec{0, 0}
	const iter = 1024
	// size := vec{7, 7}
	// end := vec{6, 6}
	// start := vec{0, 0}
	// const iter = 12
	// input := readInput(strings.NewReader(testInput))
	g := newGrid(size)
	for i, v := range input {
		if i >= iter {
			break
		}
		g.set(v, true)
	}
	// g.print(os.Stdout)

	var result vec
	for i := iter + 1; i < len(input); i++ {
		g.set(input[i], true)
		distances := dijkstra(g, start)
		if distances[end] == 0 {
			result = input[i]
			break
		}
	}
	fmt.Println(result)
}

func puzzle1() {
	size := vec{71, 71}
	end := vec{70, 70}
	start := vec{0, 0}
	const iter = 1024
	g := newGrid(size)
	// input := readInput(strings.NewReader(testInput))
	// fmt.Println(input)
	for i, v := range input {
		if i >= iter {
			break
		}
		g.set(v, true)
	}
	// g.print(os.Stdout)

	result := dijkstra(g, start)
	fmt.Println(result[end])
}

type (
	grid [][]bool
	vec  struct{ x, y int }
)

func newGrid(size vec) grid {
	g := make([][]bool, size.y)
	for yi := 0; yi < size.y; yi++ {
		g[yi] = make([]bool, size.x)
	}
	return g
}

func (g grid) set(p vec, val bool) {
	g[p.y][p.x] = val
}

func (g grid) get(p vec) bool {
	return g[p.y][p.x]
}

func (g grid) width() int {
	return len(g[0])
}

func (g grid) height() int {
	return len(g)
}

func (g grid) print(w io.Writer) {
	for _, row := range g {
		for _, blocked := range row {
			if blocked {
				w.Write([]byte("#"))
			} else {
				w.Write([]byte("."))
			}
		}
		w.Write([]byte("\n"))
	}
}

func dijkstra(g grid, start vec) map[vec]int {
	s := &gridSearcher{
		distances: map[vec]int{},
		horizon:   map[vec]int{},
	}
	return s.dijkstra(start, g)
}

type gridSearcher struct {
	distances map[vec]int
	horizon   map[vec]int
}

func (s *gridSearcher) pop() (vec, int) {
	type vecDist struct {
		p vec
		d int
	}
	sorted := slices.SortedFunc(func(yield func(vecDist) bool) {
		for vec, dist := range s.horizon {
			if !yield(vecDist{vec, dist}) {
				return
			}
		}
	}, func(a, b vecDist) int { return a.d - b.d })
	p, d := sorted[0].p, sorted[0].d
	delete(s.horizon, p)
	s.distances[p] = d
	return p, d
}

func (s *gridSearcher) put(v vec, d int) {
	existingDist, exists := s.horizon[v]
	if !exists || existingDist > d {
		if v == (vec{5, 4}) {
		}
		s.horizon[v] = d
	}
}

func (s *gridSearcher) dijkstra(start vec, g grid) map[vec]int {
	s.horizon[start] = 0
	for len(s.horizon) > 0 {
		current, d := s.pop()
		s._dijkstra(current, d, g)
	}
	return s.distances
}

func (s *gridSearcher) _dijkstra(current vec, distance int, g grid) {
	if current.x > 0 {
		left := vec{current.x - 1, current.y}
		if g.get(left) {
			goto right
		}
		if _, visited := s.distances[left]; !visited {
			s.put(left, distance+1)
		}
	}
right:
	if current.x < g.width()-1 {
		right := vec{current.x + 1, current.y}
		if g.get(right) {
			goto top
		}
		if _, visited := s.distances[right]; !visited {
			s.put(right, distance+1)
		}
	}
top:
	if current.y > 0 {
		top := vec{current.x, current.y - 1}
		if g.get(top) {
			goto bottom
		}
		if _, visited := s.distances[top]; !visited {
			s.put(top, distance+1)
		}
	}
bottom:
	if current.y < g.height()-1 {
		bottom := vec{current.x, current.y + 1}
		if g.get(bottom) {
			return
		}
		if _, visited := s.distances[bottom]; !visited {
			s.put(bottom, distance+1)
		}
	}
}
