package main

import "testing"

var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day22
// cpu: Apple M1 Pro
// BenchmarkPuzzle1-10    	      37	  31566277 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPuzzle1(b *testing.B) {
	var result int
	for n := 0; n < b.N; n++ {
		result = count2000Prices(nums)
	}
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day22
// cpu: Apple M1 Pro
// BenchmarkPuzzle2-10    	       2	 581697021 ns/op	804252748 B/op	  173871 allocs/op
func BenchmarkPuzzle2(b *testing.B) {
	var result int
	for n := 0; n < b.N; n++ {
		result = findMaxBananas(nums)
	}
	anchor = result
}
