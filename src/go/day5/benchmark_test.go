package main

import "testing"

var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day5
// cpu: Apple M1 Pro
// BenchmarkCountGood-10    	    4700	    247612 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCountGood(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = sum(arrays, goodCounter(rules))
	}
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day5
// cpu: Apple M1 Pro
// BenchmarkCountBad-10    	    3920	    357964 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCountBad(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = sum(arrays, badCounter(rules))
	}
	anchor = result
}
