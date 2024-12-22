package main

import "testing"

var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day21
// cpu: Apple M1 Pro
// BenchmarkPuzzle1-10    	   53643	     21704 ns/op	   40209 B/op	     317 allocs/op
func BenchmarkPuzzle1(b *testing.B) {
	var result int
	for n := 0; n < b.N; n++ {
		countComplexities2(codes, 2)
	}
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day21
// cpu: Apple M1 Pro
// BenchmarkPuzzle2-10    	    5016	    219871 ns/op	  189354 B/op	    4221 allocs/op
func BenchmarkPuzzle2(b *testing.B) {
	var result int
	for n := 0; n < b.N; n++ {
		countComplexities2(codes, 25)
	}
	anchor = result
}
