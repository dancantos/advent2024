package main

import (
	"regexp"
	"strconv"
)

var mulregex = regexp.MustCompile(`mul\(([1-9][0-9]{0,2}),([1-9][0-9]{0,2})\)`)

func mulall(input string) int {
	matches := mulregex.FindAllStringSubmatch(input, -1)
	sum := 0
	for _, match := range matches {
		n, _ := strconv.Atoi(string(match[1]))
		m, _ := strconv.Atoi(string(match[2]))
		sum += n * m
	}
	return sum
}
