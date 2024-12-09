package main

import "testing"

var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day9
// cpu: Apple M1 Pro
// BenchmarkPuzzle2-10    	       3	 375849208 ns/op	 4103130 B/op	      30 allocs/op
func BenchmarkPuzzle2(b *testing.B) {
	var result int
	for n := 0; n < b.N; n++ {
		p := process(input)
		l := left2(p)
		checksum(l)
	}
	anchor = result
}
