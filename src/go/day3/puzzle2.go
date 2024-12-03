package main

import (
	"regexp"
	"strconv"
)

var instructionRegex = regexp.MustCompile(`(?:mul\(([1-9][0-9]{0,2}),([1-9][0-9]{0,2})\)|do\(\)|don't\(\))`)

func mulWithInstruction(input string) int {
	matches := instructionRegex.FindAllStringSubmatch(input, -1)
	sum := 0
	on := true
	for _, match := range matches {
		switch match[0] {
		case "do()":
			on = true
			continue
		case "don't()":
			on = false
			continue
		}
		if !on {
			continue
		}
		n, _ := strconv.Atoi(string(match[1]))
		m, _ := strconv.Atoi(string(match[2]))
		sum += n * m
	}
	return sum
}
