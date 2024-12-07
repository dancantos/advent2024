package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/dancantos/advent2024/src/go/pkg/must"
	"github.com/dancantos/advent2024/src/go/pkg/timeit"
)

//go:embed input.txt
var input []byte
var equations = readInput(bytes.NewReader(input))

func readInput(r io.Reader) []equation {
	result := make([]equation, 0)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		equationStrings := strings.Split(s.Text(), " ")
		equation := equation{
			target: must.Return(strconv.Atoi(equationStrings[0][:len(equationStrings[0])-1])),
			terms:  make([]int, len(equationStrings)-1),
		}
		for i := 1; i < len(equationStrings); i++ {
			equation.terms[i-1] = must.Return(strconv.Atoi(equationStrings[i]))
		}
		result = append(result, equation)
	}
	return result
}

type equation struct {
	target int
	terms  []int
}

func main() {
	// equations := readInput(strings.NewReader(testInput))
	fmt.Printf("puzzle 1: solveable targets (+*): %d\n", countSolveableTargets(equations, isSolveable))
	for range timeit.Run(1) {
		countSolveableTargets(equations, isSolveable)
	}

	fmt.Println()

	fmt.Printf("puzzle 2: solveable targets (+*||): %d\n", countSolveableTargets(equations, isSolveableWithConcat))
	for range timeit.Run(1) {
		countSolveableTargets(equations, isSolveableWithConcat)
	}
}

func countSolveableTargets(equations []equation, filter func(equation) bool) int {
	count := 0
	for _, eq := range equations {
		if filter(eq) {
			count += eq.target
		}
	}
	return count
}
