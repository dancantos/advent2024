package main

import "testing"

var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day11
// cpu: Apple M1 Pro
// Benchmark25Blinks-10    	    2086	    604618 ns/op	  449474 B/op	    3474 allocs/op
func Benchmark25Blinks(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = fullCount(arr, 25)
	}
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day11
// cpu: Apple M1 Pro
// Benchmark75Blinks-10    	      49	  23806292 ns/op	 8884101 B/op	   33626 allocs/op
func Benchmark75Blinks(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = fullCount(arr, 75)
	}
	anchor = result
}
