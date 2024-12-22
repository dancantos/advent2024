package main

import "testing"

var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day21
// cpu: Apple M1 Pro
// BenchmarkPuzzle1-10    	  167002	      7181 ns/op	   11470 B/op	     108 allocs/op
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
// BenchmarkPuzzle2-10    	   26112	     46083 ns/op	   41061 B/op	     889 allocs/op
func BenchmarkPuzzle2(b *testing.B) {
	var result int
	for n := 0; n < b.N; n++ {
		sumComplexities(codes, nums, 25)
	}
	anchor = result
}
