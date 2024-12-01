package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"strings"
)

var (
	a1 = randArr()
	a2 = randArr()
)

func randArr() []int {
	arr := make([]int, 1e4)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(1e4)
	}
	return arr
}

func main() {
	a1, a2, err := readInput(data)
	if err != nil {
		panic(err)
	}
	// fmt.Println(a1, a2)
	// fmt.Println(sumdiffsnaive(a1, a2))
	fmt.Println(sumdiffsdoesthiswork(a1, a2))
}

//go:embed input.txt
var data []byte

func readInput(data []byte) ([]int, []int, error) {
	a1 := make([]int, 0)
	a2 := make([]int, 0)

	s := bufio.NewScanner(bytes.NewReader(data))
	s.Split(bufio.ScanLines)
	count := 0
	for s.Scan() {
		cut1, cut2, found := bytes.Cut(s.Bytes(), []byte(" "))
		if !found {
			return nil, nil, fmt.Errorf("malformed input (line:%d)", count)
		}
		n1, err := strconv.Atoi(string(cut1))
		if err != nil {
			return nil, nil, fmt.Errorf("malformed input (line:%d)", count)
		}

		n2, err := strconv.Atoi(strings.TrimSpace(string(cut2)))
		if err != nil {
			return nil, nil, fmt.Errorf("malformed input (line:%d)", count)
		}

		a1 = append(a1, n1)
		a2 = append(a2, n2)
		count++
	}
	return a1, a2, nil
}

func sumdiffsnaive(a1, a2 []int) int {
	slices.Sort(a1)
	slices.Sort(a2)

	sum := 0
	for i := 0; i < len(a1); i++ {
		if a1[i] < a2[i] {
			sum += a2[i] - a1[i]
			continue
		}
		sum += a1[i] - a2[i]
	}
	return sum
}

// This algorithm is the same as the naive algorithm but ignores the sort
// This is based on the observation that altering any 2 elements in either list by the same amount does not affect the resulting sum of diffs.
func sumdiffsdoesthiswork(a1, a2 []int) int {
	sum := 0
	for i := 0; i < len(a1); i++ {
		if a1[i] < a2[i] {
			sum += a2[i] - a1[i]
			continue
		}
		sum += a1[i] - a2[i]
	}
	return sum
}
