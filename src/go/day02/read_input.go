package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"io"
	"strconv"
	"strings"
)

//go:embed input.txt
var data []byte

var Input = func(data []byte) [][]int {
	input, err := readInput(bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	return input
}(data)

func readInput(data io.Reader) ([][]int, error) {
	result := make([][]int, 0)
	s := bufio.NewScanner(data)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		if len(s.Text()) == 0 {
			continue
		}
		numStrings := strings.Split(s.Text(), " ")
		nums := make([]int, len(numStrings))
		for i, val := range numStrings {
			n, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			nums[i] = n
		}
		result = append(result, nums)
	}
	return result, nil
}
