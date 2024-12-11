package main

import "testing"

var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day11
// cpu: Apple M1 Pro
// Benchmark25Blinks-10    	    3016	    368452 ns/op	  503662 B/op	    1651 allocs/op
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
// Benchmark75Blinks-10    	     100	  10469668 ns/op	 3174507 B/op	    4040 allocs/op
func Benchmark75Blinks(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = fullCount(arr, 75)
	}
	anchor = result
}
