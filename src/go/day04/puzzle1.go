package main

func countXmas(arr [][]rune) int {
	m := len(arr)
	n := len(arr[0])
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j] == 'X' || arr[i][j] == 'S' {
				count += detectXmas(arr, m, n, i, j)
			}
		}
	}
	return count
}

func detectXmas(arr [][]rune, m, n, i, j int) int {
	count := 0
	// Left Down
	if i < m-3 && j > 2 &&
		(arr[i][j] == 'X' && arr[i+1][j-1] == 'M' && arr[i+2][j-2] == 'A' && arr[i+3][j-3] == 'S' ||
			arr[i][j] == 'S' && arr[i+1][j-1] == 'A' && arr[i+2][j-2] == 'M' && arr[i+3][j-3] == 'X') {
		count++
	}

	// Down
	if i < m-3 &&
		(arr[i][j] == 'X' && arr[i+1][j] == 'M' && arr[i+2][j] == 'A' && arr[i+3][j] == 'S' ||
			arr[i][j] == 'S' && arr[i+1][j] == 'A' && arr[i+2][j] == 'M' && arr[i+3][j] == 'X') {
		count++
	}

	// Down Right
	if i < m-3 && j < n-3 &&
		(arr[i][j] == 'X' && arr[i+1][j+1] == 'M' && arr[i+2][j+2] == 'A' && arr[i+3][j+3] == 'S' ||
			arr[i][j] == 'S' && arr[i+1][j+1] == 'A' && arr[i+2][j+2] == 'M' && arr[i+3][j+3] == 'X') {
		count++
	}

	// Right
	if j < n-3 &&
		(arr[i][j] == 'X' && arr[i][j+1] == 'M' && arr[i][j+2] == 'A' && arr[i][j+3] == 'S' ||
			arr[i][j] == 'S' && arr[i][j+1] == 'A' && arr[i][j+2] == 'M' && arr[i][j+3] == 'X') {
		count++
	}

	return count
}
