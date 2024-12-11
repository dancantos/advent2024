package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/dancantos/advent2024/src/go/pkg/must"
	"github.com/dancantos/advent2024/src/go/pkg/timeit"
)

//go:embed input.txt
var input string
var arr = readInput(input)

func readInput(s string) []uint64 {
	split := strings.Split(input, " ")
	result := make([]uint64, len(split))
	for i, n := range strings.Split(input, " ") {
		result[i] = uint64(must.Return(strconv.Atoi(n)))
	}
	return result
}

func main() {
	fmt.Println("Puzzle1 (25 blinks):", fullCount(arr, 25))
	for range timeit.Run(1) {
		fullCount(arr, 25)
	}
	fmt.Println("Puzzle2 (75 blinks):", fullCount(arr, 75))
	for range timeit.Run(1) {
		fullCount(arr, 75)
	}
}

func fullCount(arr []uint64, iterations int) int {
	count := 0
	m := make(memory)
	for _, n := range arr {
		count += countStones(n, iterations, m)
	}
	return count
}

func countStones(n uint64, depth int, m memory) int {
	if result, exists := m.get(n, depth); exists {
		return result
	}
	if depth == 0 {
		return 1
	}
	result := 0
	switch {
	case n == 0:
		result = countStones(1, depth-1, m)
	case evenDigits(n):
		l, r := splitIntByMath(n)
		result = countStones(l, depth-1, m) + countStones(r, depth-1, m)
	default:
		result = countStones(n*2024, depth-1, m)
	}
	m.record(n, depth, result)
	return result
}

func evenDigits(n uint64) bool {
	for n >= 100 {
		n /= 100
	}
	return n >= 10
}

func splitIntByMath(n uint64) (left, right uint64) {
	digits := uint64(math.Ceil(math.Log10(float64(n) + 0.1)))
	exp := uint64(math.Pow10(int(digits) / 2))
	return n / exp, n % exp
}

type memory map[uint64]map[int]int

func (m memory) record(n uint64, depth, result int) {
	r, exists := m[n]
	if !exists {
		r = make(map[int]int)
		m[n] = r
	}
	// fmt.Println(r, depth)
	r[depth] = result
}

func (m memory) get(n uint64, depth int) (int, bool) {
	if r, exists := m[n]; exists {
		result, exists := r[depth]
		return result, exists
	}
	return 0, false
}
