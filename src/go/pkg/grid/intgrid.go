package grid

import (
	"bufio"
	"io"

	"github.com/dancantos/advent2024/src/go/pkg/lin"
)

type IntGrid struct {
	size lin.Vec[int]
	grid [][]int
}

func ReadIntGrid(r io.Reader, callback func(lin.Vec[int], rune)) IntGrid {
	g := IntGrid{
		grid: make([][]int, 0),
	}
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	y := 0
	for s.Scan() {
		g.size.X = len(s.Text())
		row := make([]int, len(s.Text()))
		for x, char := range s.Text() {
			row[x] = int(char - '0')
			callback(lin.Vec[int]{x, y}, char)
		}
		g.grid = append(g.grid, row)
		y++
	}
	return g
}
