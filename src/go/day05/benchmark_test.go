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
// BenchmarkCountBad-10    	    2486	    485612 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCountBad(b *testing.B) {
	var result int

	var cp [][]int
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		cp = copyInput(arrays)
		b.StartTimer()
		result = sum(cp, badCounter(rules))
	}
	anchor = result
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dancantos/advent2024/src/go/day5
// cpu: Apple M1 Pro
// BenchmarkCountBadNoMutation-10    	    2329	    514100 ns/op	   49250 B/op	    1005 allocs/op
func BenchmarkCountBadNoMutation(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = sum(arrays, badCounterNoMutation(rules))
	}
	anchor = result
}

func copyInput(arrays [][]int) [][]int {
	result := make([][]int, len(arrays))
	for _, arr := range arrays {
		cp := make([]int, len(arr))
		copy(cp, arr)
		result = append(result, cp)
	}
	return result
}
