package main

import "testing"

var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day21
// cpu: Apple M1 Pro
// BenchmarkPuzzle1-10    	  388218	      3064 ns/op	    6056 B/op	     115 allocs/op
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
// BenchmarkPuzzle2-10    	   48628	     24130 ns/op	   40248 B/op	     896 allocs/op
func BenchmarkPuzzle2(b *testing.B) {
	var result int
	for n := 0; n < b.N; n++ {
		sumComplexities(codes, nums, 25)
	}
	anchor = result
}
