package main

import "testing"

var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day13
// cpu: Apple M1 Pro
// BenchmarkPuzzle1-10    	   25304	     41422 ns/op	   40960 B/op	    1920 allocs/op
func BenchmarkPuzzle1(b *testing.B) {
	var result int
	for n := 0; n < b.N; n++ {
		result = solve(p1)
	}
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day21
// cpu: Apple M1 Pro
// BenchmarkPuzzle2-10    	    5434	    217652 ns/op	  189358 B/op	    4221 allocs/op
func BenchmarkPuzzle2(b *testing.B) {
	var result int
	for n := 0; n < b.N; n++ {
		result = solve(p2)
	}
	anchor = result
}
