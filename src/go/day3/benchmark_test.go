package main

import "testing"

var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day3
// cpu: Apple M1 Pro
// BenchmarkMulAll-10    	    3301	    358714 ns/op	   97317 B/op	    1357 allocs/op
func BenchmarkMulAll(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = mulall(input)
	}
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day3
// cpu: Apple M1 Pro
// BenchmarkMulWithInstruction-10    	    1554	    747465 ns/op	  131533 B/op	    1502 allocs/op
func BenchmarkMulWithInstruction(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = mulWithInstruction(input)
	}
	anchor = result
}
