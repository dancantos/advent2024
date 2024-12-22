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
// BenchmarkCountMas-10    	   14001	     84844 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCountMas(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = countMas(input)
	}
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day4
// cpu: Apple M1 Pro
// BenchmarkCountMasSmartButSlow-10    	    6028	    194865 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCountMasSmartButSlow(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = _countMas(input)
	}
	anchor = result
}

func _countMas(arr [][]rune) int {
	m := len(arr)
	n := len(arr[0])
	count := 0
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if arr[i][j] == 'A' {
				count += detectMasSmartButSlow(arr, i, j)
			}
		}
	}
	return count
}
