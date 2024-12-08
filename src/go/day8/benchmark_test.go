package main

import "testing"

var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day8
// cpu: Apple M1 Pro
// BenchmarkPuzzle1-10    	   28632	     43659 ns/op	   49213 B/op	     825 allocs/op
func BenchmarkPuzzle1(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = countAntinodes(inputGrid, size, iterateAntinodes)
	}
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day8
// cpu: Apple M1 Pro
// BenchmarkPuzzle2-10    	    8224	    132405 ns/op	  118563 B/op	     860 allocs/op
func BenchmarkPuzzle2(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = countAntinodes(inputGrid, size, iterateMoreAntinodes)
	}
	anchor = result
}
