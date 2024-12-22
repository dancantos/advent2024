package main

import "testing"

var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day21
// cpu: Apple M1 Pro
// BenchmarkPuzzle1-10    	  500146	      2458 ns/op	    4560 B/op	      84 allocs/op
func BenchmarkPuzzle1(b *testing.B) {
	var result int
	for n := 0; n < b.N; n++ {
		sumComplexities(codes, nums, 2)
	}
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day21
// cpu: Apple M1 Pro
// BenchmarkPuzzle2-10    	   51781	     22882 ns/op	   33000 B/op	     867 allocs/op
func BenchmarkPuzzle2(b *testing.B) {
	var result int
	for n := 0; n < b.N; n++ {
		sumComplexities(codes, nums, 25)
	}
	anchor = result
}
