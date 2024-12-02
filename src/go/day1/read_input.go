package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var data []byte

var A1, A2 = func() ([]int, []int) {
	a1, a2, err := readInput(data)
	if err != nil {
		panic(err)
	}
	return a1, a2
}()

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
