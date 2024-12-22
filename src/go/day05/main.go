package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"

	_ "embed"

	"github.com/dancantos/advent2024/src/go/pkg/must"
	"github.com/dancantos/advent2024/src/go/pkg/timeit"
)

//go:embed input.txt
var input []byte

var rules, arrays = readInput(bytes.NewReader(input))

func main() {
	fmt.Printf("Puzzle 1 (Count Good): %d\n", sum(arrays, goodCounter(rules)))
	for range timeit.Run(1) {
		sum(arrays, goodCounter(rules))
	}

	fmt.Println()

	fmt.Printf("Puzzle 2 (Count Bad Sorted): %d\n", sum(arrays, badCounter(rules)))
	for range timeit.Run(1) {
		sum(arrays, badCounter(rules))
	}
}

func sum(lines [][]int, counter func([]int) int) int {
	sum := 0
	for _, l := range lines {
		sum += counter(l)
	}
	return sum
}

func isSorted(arr []int, rules map[int][]int) bool {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if slices.Contains(rules[arr[j]], arr[i]) {
				return false
			}
		}
	}
	return true
}

func sort(rules map[int][]int) func(a, b int) int {
	return func(a, b int) int {
		if a == b {
			return 0
		}
		if _, ok := rules[b]; !ok {
			return 0
		}
		if slices.Contains(rules[b], a) {
			return -1
		}
		return 1
	}
}

func readInput(r io.Reader) (map[int][]int, [][]int) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	section := 0
	rules := map[int][]int{}
	lines := [][]int{}
	for s.Scan() {
		if s.Text() == "" {
			section++
			continue
		}
		if section == 0 {
			split := strings.Split(s.Text(), "|")
			k, v := must.Return(strconv.Atoi(split[0])), must.Return(strconv.Atoi(split[1]))
			rules[k] = append(rules[k], v)
		}
		if section == 1 {
			lines = append(lines, slices.Collect(func(yield func(int) bool) {
				for _, str := range strings.Split(s.Text(), ",") {
					yield(must.Return(strconv.Atoi(str)))
				}
			}))
		}
	}
	return rules, lines
}
