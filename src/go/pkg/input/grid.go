package input

import (
	"bufio"
	"io"
)

func ReadGrid(r io.Reader, callback func(x, y int, char rune)) (width, height int) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		width = len(s.Text())
		for x, char := range s.Text() {
			callback(x, height, char)
		}
		height++
	}
	return
}
