package main

import "fmt"

func puzzle2() {
	fmt.Println(similarityScore(a1, a2))
}

func similarityScore(a1, a2 []int) int {
	a2Counts := make(map[int]int, len(a1))
	for i := 0; i < len(a1); i++ {
		a2Counts[a2[i]]++
	}
	sum := 0
	for _, n := range a1 {
		sum += n * a2Counts[n]
	}
	return sum
}
