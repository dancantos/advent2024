package grid

import (
	"bufio"
	"io"

	"github.com/dancantos/advent2024/src/go/pkg/lin"
)

type Grid[T any] struct {
	grid          [][]T
	Width, Height int
}

func ReadGrid[T any](r io.Reader, cell func(rune) T) Grid[T] {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	result := Grid[T]{
		grid: make([][]T, 0),
	}
	for s.Scan() {
		result.Width = len(s.Text())
		row := make([]T, len(s.Text()))
		for i, char := range s.Text() {
			row[i] = cell(char)
		}
		result.grid = append(result.grid, row)
		result.Height++
	}
	return result
}

func ReadInt(r rune) int {
	return int(r - '0')
}

func ReadChar(r rune) rune {
	return r
}

func (g Grid[T]) Get(p lin.IVec) T {
	return g.grid[p.X][p.Y]
}

func (g Grid[T]) Set(p lin.IVec, val T) {
	g.grid[p.X][p.Y] = val
}

func (g Grid[T]) In(p lin.IVec) bool {
	return 0 <= p.X && p.X < g.Width && 0 <= p.Y && p.Y < g.Height
}
