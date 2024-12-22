package main

import "slices"

func countMas(arr [][]rune) int {
	m := len(arr)
	n := len(arr[0])
	count := 0
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if arr[i][j] == 'A' {
				count += detectMas(arr, i, j)
				// count += detectMasSmartButSlow(arr, i, j)
			}
		}
	}
	return count
}

func detectMas(arr [][]rune, i, j int) int {
	if ((arr[i-1][j-1] == 'M' && arr[i+1][j+1] == 'S') || (arr[i-1][j-1] == 'S' && arr[i+1][j+1] == 'M')) &&
		((arr[i-1][j+1] == 'M' && arr[i+1][j-1] == 'S') || (arr[i-1][j+1] == 'S' && arr[i+1][j-1] == 'M')) {
		return 1
	}
	return 0
}

var possible = []string{"MMSS", "MSMS", "SMSM", "SSMM"}

func detectMasSmartButSlow(arr [][]rune, i, j int) int {
	testStr := string([]rune{arr[i-1][j-1], arr[i-1][j+1], arr[i+1][j-1], arr[i+1][j+1]})
	if slices.Contains(possible, testStr) {
		return 1
	}
	return 0
}
