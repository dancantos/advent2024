package main

import "testing"

var anchor int

func BenchmarkPuzzle1(b *testing.B) {
	var result int
	for n := 0; n < b.N; n++ {
		result = solve(p1)
	}
	anchor = result
}

func BenchmarkPuzzle2(b *testing.B) {
	var result int
	for n := 0; n < b.N; n++ {
		result = solve(p2)
	}
	anchor = result
}
