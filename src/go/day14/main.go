package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"iter"
	"os"
	"regexp"
	"strconv"

	"github.com/dancantos/advent2024/src/go/pkg/it"
	"github.com/dancantos/advent2024/src/go/pkg/must"
)

//go:embed input.txt
var input []byte

var robots = readInput(bytes.NewReader(input))

var stateRegex = regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

func readInput(r io.Reader) []state {
	result := []state{}
	for line := range it.ReadLines(r) {
		match := stateRegex.FindStringSubmatch(line)
		if match == nil {
			panic("RUH ROH")
		}
		result = append(result, state{
			p: vec{must.Return(strconv.Atoi(match[1])), must.Return(strconv.Atoi(match[2]))},
			v: vec{must.Return(strconv.Atoi(match[3])), must.Return(strconv.Atoi(match[4]))},
		})
	}
	return result
}

func main() {
	// const sampleInput = `p=0,4 v=3,-3
	// p=6,3 v=-1,-3
	// p=10,3 v=-1,2
	// p=2,0 v=2,-1
	// p=0,0 v=1,3
	// p=3,0 v=-2,-2
	// p=7,6 v=-1,-3
	// p=3,0 v=-1,-2
	// p=9,3 v=2,3
	// p=7,3 v=-1,2
	// p=2,4 v=2,-3
	// p=9,5 v=-3,-3`

	// robots := readInput(strings.NewReader(sampleInput))
	g := newGrid(SPACE_WIDTH, SPACE_HEIGHT)
	f, err := os.CreateTemp(".", "robots")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// display loop
	time := 0
	for count := 0; count < SPACE_WIDTH*SPACE_HEIGHT; count++ {
		time++
		for _, r := range robots {
			g.clear(r.p.x, r.p.y)
		}
		for i, r := range robots {
			r = r.next()
			robots[i] = r
			g.increment(r.p.x, r.p.y)
		}
		fmt.Fprintf(f, "TIME: %ds\n", time)
		g.print(f)
	}
}

func puzzle1() {
	g := newGrid(SPACE_WIDTH, SPACE_HEIGHT)
	const iterCount = 100
	for _, r := range robots {
		for n := 0; n < iterCount; n++ {
			r = r.next()
		}
		g.increment(r.p.x, r.p.y)
	}

	mul := 1
	for _, quadrant := range g.quadrants() {
		sum := 0
		for x, y := range quadrant.cells() {
			sum += g.get(x, y)
		}
		if sum != 0 {
			mul *= sum
		}
	}
	fmt.Println(mul)
}

const (
	SPACE_WIDTH  = 101
	SPACE_HEIGHT = 103
	// SPACE_WIDTH  = 11
	// SPACE_HEIGHT = 7
)

type (
	vec   struct{ x, y int }
	state struct {
		p, v vec
	}
	grid     [][]int
	quadrant struct{ minX, maxX, minY, maxY int }
)

func newGrid(width, height int) grid {
	g := make(grid, height)
	for i := 0; i < height; i++ {
		g[i] = make([]int, width)
	}
	return g
}

func (g grid) quadrants() []quadrant {
	width, height := len(g[0]), len(g)
	return []quadrant{
		{0, width/2 - 1, 0, height/2 - 1},
		{0, width/2 - 1, height/2 + 1, height - 1},
		{width/2 + 1, width - 1, 0, height/2 - 1},
		{width/2 + 1, width - 1, height/2 + 1, height - 1},
	}
}

func (g grid) get(x, y int) int {
	return g[y][x]
}

func (g grid) increment(x, y int) {
	g[y][x]++
}

func (g grid) clear(x, y int) {
	g[y][x] = 0
}

func (g grid) print(out io.Writer) {
	for _, row := range g {
		for _, entry := range row {
			if entry == 0 {
				fmt.Fprint(out, ".")
			} else {
				fmt.Fprintf(out, "%d", entry)
			}
		}
		fmt.Fprintln(out)
	}
}

func (q quadrant) cells() iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for x := q.minX; x <= q.maxX; x++ {
			for y := q.minY; y <= q.maxY; y++ {
				if !yield(x, y) {
					return
				}
			}
		}
	}
}

func (s state) next() state {
	next := state{
		p: s.p.add(s.v),
		v: s.v,
	}
	// teleport
	for next.p.x < 0 {
		next.p.x += SPACE_WIDTH
	}
	for next.p.x >= SPACE_WIDTH {
		next.p.x -= SPACE_WIDTH
	}
	for next.p.y < 0 {
		next.p.y += SPACE_HEIGHT
	}
	for next.p.y >= SPACE_HEIGHT {
		next.p.y -= SPACE_HEIGHT
	}
	next.p.y %= SPACE_HEIGHT

	if next.p.x < 0 || next.p.y < 0 {
		fmt.Println("RUH ROH")
	}
	return next
}

// // in place state update without need for copy
// func (s *state) inext() {
// 	s.p = s.p.add(s.v)
// 	// teleport
// 	s.p.x %= SPACE_WIDTH
// 	s.p.y %= SPACE_HEIGHT
// }

func (u vec) add(v vec) vec {
	return vec{u.x + v.x, u.y + v.y}
}
