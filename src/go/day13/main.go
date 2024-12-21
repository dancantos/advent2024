package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/dancantos/advent2024/src/go/pkg/lin"
	"github.com/dancantos/advent2024/src/go/pkg/must"
)

//go:embed input.txt
var input string

const offset = 10000000000000

var (
	p1 = readInput(input)

	// copy p1 with offset added to x and y
	p2 = func() []Problem {
		result := make([]Problem, 0, len(p1))
		for _, p := range p1 {
			result = append(result, Problem{
				m:      p.m,
				target: []float64{p.target[0] + offset, p.target[1] + offset},
			})
		}
		return result
	}()
)

func main() {
	const TOLERANCE float64 = 1e-3

	fmt.Printf("Puzzle 1 (small arcade): %d\n", solve(p1))
	fmt.Printf("Puzzle 2 (big arcade): %d", solve(p2))
}

func readInput(s string) []Problem {
	matches := problemRegex.FindAllStringSubmatch(s, -1)
	result := make([]Problem, len(matches))
	for i, match := range matches {
		mat := lin.Mat[float64]{
			{float64(must.Return(strconv.Atoi(match[1]))), float64(must.Return(strconv.Atoi(match[3])))},
			{float64(must.Return(strconv.Atoi(match[2]))), float64(must.Return(strconv.Atoi(match[4])))},
		}
		target := []float64{float64(must.Return(strconv.Atoi(match[5]))), float64(must.Return(strconv.Atoi(match[6])))}
		result[i] = Problem{mat, target}
	}
	return result
}

var problemRegex = regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)\nButton B: X\+(\d+), Y\+(\d+)\nPrize: X=(\d+), Y=(\d+)`)

type Problem struct {
	m      lin.Mat[float64]
	target []float64
}

func solve(input []Problem) int {
	const TOLERANCE float64 = 1e-3
	sum := 0
	for _, p := range input {
		solved := p.m.LUSolve(p.target)
		solution := roundOrOmit(solved, TOLERANCE)
		if solution == nil {
			continue
		}
		sum += 3*solution[0] + solution[1]
	}
	return sum
}

func roundOrOmit(solution []float64, tolerance float64) []int {
	result := make([]int, len(solution))
	for i, val := range solution {
		r := math.Round(val)
		diff := abs(r - val)
		if diff > tolerance {
			return nil
		}
		result[i] = int(r)
	}
	return result
}

func abs(f float64) float64 {
	if f < 0 {
		f = -f
	}
	return f
}
