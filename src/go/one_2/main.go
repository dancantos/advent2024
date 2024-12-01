package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

var (
	a1 = randArr()
	a2 = randArr()
)

func randArr() []int {
	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(10)
	}
	return arr
}

func main() {
	a1, a2, err := readInput(data)
	if err != nil {
		panic(err)
	}
	// fmt.Println(a1, a2)
	fmt.Println(similarityScore(a1, a2))
}

func similarityScore(a1, a2 []int) int {
	a2Counts := make(map[int]int)
	for i := 0; i < len(a1); i++ {
		a2Counts[a2[i]]++
	}
	// fmt.Println(a2Counts)
	sum := 0
	for _, n := range a1 {
		sum += n * a2Counts[n]
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
