package main

import "testing"

var anchor int

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day4
// cpu: Apple M1 Pro
// BenchmarkCountXmas-10    	    5540	    226889 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCountXmas(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = countXmas(input)
	}
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day4
// cpu: Apple M1 Pro
// BenchmarkCountMas-10    	   14115	     87977 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCountMas(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = countMas(input)
	}
	anchor = result
}
