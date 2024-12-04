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
var data []byte

func main() {
	arr := readInput(bytes.NewReader(data))
	fmt.Printf("Puzzle 1 (Count \"XMAS\"): %d\n", countXmas(arr))
	for range timeit.Run(1) {
		countXmas(arr)
	}

	fmt.Println()

	fmt.Printf("Puzzle 2 (Count \"MAS\"): %d\n", countMas(arr))
	for range timeit.Run(1) {
		countMas(arr)
	}
}

func readInput(r io.Reader) [][]rune {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	arr := make([][]rune, 0)
	for s.Scan() {
		arr = append(arr, []rune(s.Text()))
	}
	return arr
}
