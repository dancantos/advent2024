package main

import "testing"

var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day8
// cpu: Apple M1 Pro
// BenchmarkPuzzle1-10    	   55172	     21884 ns/op	   36232 B/op	     794 allocs/op
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
// BenchmarkPuzzle2-10    	   34711	     34846 ns/op	   40456 B/op	     794 allocs/op
func BenchmarkPuzzle2(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = countAntinodes(inputGrid, size, iterateMoreAntinodes)
	}
	anchor = result
}
