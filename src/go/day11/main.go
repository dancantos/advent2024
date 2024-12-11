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

func readInput(s string) []int {
	split := strings.Split(input, " ")
	result := make([]int, len(split))
	for i, n := range strings.Split(input, " ") {
		result[i] = int(must.Return(strconv.Atoi(n)))
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

func fullCount(arr []int, iterations int) int {
	count := 0
	m := newMemory(iterations)
	for _, n := range arr {
		count += countStones(n, iterations, m)
	}
	return count
}

func countStones(n int, depth int, m memory) int {
	if depth == 0 {
		return 1
	}
	if result, exists := m.get(n, depth); exists {
		return result
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

func evenDigits(n int) bool {
	for n >= 100 {
		n /= 100
	}
	return n >= 10
}

func splitIntByMath(n int) (left, right int) {
	digits := int(math.Ceil(math.Log10(float64(n) + 0.1)))
	exp := int(math.Pow10(int(digits) / 2))
	return n / exp, n % exp
}

type memory struct {
	maxDepth int
	m        map[int][]int
}

func newMemory(maxDepth int) memory {
	return memory{
		maxDepth: maxDepth,
		m:        make(map[int][]int),
	}
}

func (m memory) record(n int, depth, result int) {
	r, exists := m.m[n]
	if !exists {
		r = make([]int, m.maxDepth)
		m.m[n] = r
	}
	// fmt.Println(r, depth)
	r[depth-1] = result
}

func (m memory) get(n int, depth int) (int, bool) {
	if r, exists := m.m[n]; exists {
		result := r[depth-1]
		return result, result > 0
	}
	return 0, false
}
