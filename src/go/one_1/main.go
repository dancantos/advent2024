package main

import (
	"bufio"
	"bytes"
	"crypto/rand"
	_ "embed"
	"fmt"
	"math/big"
	"slices"
	"strconv"
	"strings"
)

// 1 2
// 4 3

// 1 + 1 = 2

// 4 2
// 1 3

// 2 + 2 = 4

var (
	size int = 5
	a1       = randArr()
	a2       = randArr()
	max      = big.NewInt(10)
)

func randArr() []int {
	arr := make([]int, size)
	for i := 0; i < len(arr); i++ {
		n, _ := rand.Int(rand.Reader, max)
		arr[i] = int(n.Int64())
	}
	return arr
}

func main() {
	a1, a2, err := readInput(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(sumdiffs(a1, a2))
}

func sumdiffs(a1, a2 []int) int {
	// a1 = slices.Clone(a1)
	// a2 = slices.Clone(a2)
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
