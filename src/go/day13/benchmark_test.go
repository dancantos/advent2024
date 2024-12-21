package main

import "testing"

var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day13
// cpu: Apple M1 Pro
// BenchmarkPuzzle1-10    	   30721	     39461 ns/op	   40960 B/op	    1920 allocs/op
func BenchmarkPuzzle1(b *testing.B) {
	var result int
	for n := 0; n < b.N; n++ {
		result = solve(p1)
	}
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day13
// cpu: Apple M1 Pro
// BenchmarkPuzzle2-10    	   29456	     39696 ns/op	   40960 B/op	    1920 allocs/op
func BenchmarkPuzzle2(b *testing.B) {
	var result int
	for n := 0; n < b.N; n++ {
		result = solve(p2)
	}
	anchor = result
}
