package main

import "testing"

var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day23
// cpu: Apple M1 Pro
// BenchmarkP1-10    	       2	 510357916 ns/op	749805080 B/op	21866602 allocs/op
func BenchmarkP1(b *testing.B) {
	var result int
	for n := 0; n < b.N; n++ {
		result = puzzle1(data)
	}
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day23
// cpu: Apple M1 Pro
// BenchmarkP2-10    	       1	5831600417 ns/op	7938325384 B/op	238892756 allocs/op
func BenchmarkP2(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		result = puzzle2(data)
	}
	anchor = len(result)
}
