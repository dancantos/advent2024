package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"io"
	"regexp"
	"strconv"

	"github.com/dancantos/advent2024/src/go/pkg/must"
)

//go:embed input.txt
var input []byte

var codes, nums = readInput(bytes.NewReader(input))

var numRegex = regexp.MustCompile(`\d+`)

func readInput(r io.Reader) ([]string, []int) {
	codes := []string{}
	nums := []int{}
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		codes = append(codes, s.Text())
		n := numRegex.FindString(s.Text())
		nums = append(nums, must.Return(strconv.Atoi(n)))
	}
	return codes, nums
}
