package main

import "testing"

var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day7
// cpu: Apple M1 Pro
// BenchmarkPuzzle1-10    	   18835	     63733 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPuzzle1(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = countSolveableTargets(equations, isSolveable)
	}
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day7
// cpu: Apple M1 Pro
// BenchmarkPuzzle2-10    	    2325	    514181 ns/op	   88994 B/op	   12025 allocs/op
func BenchmarkPuzzle2(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = countSolveableTargets(equations, isSolveableWithConcat)
	}
	anchor = result
}
