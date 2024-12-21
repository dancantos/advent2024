package lin

import "golang.org/x/exp/constraints"

type Mat[N constraints.Integer] struct {
	grid [][]N
}

func NewMat[N constraints.Integer](r, c int) Mat[N] {
	m := Mat[N]{}
	m.grid = make([][]N, r)
	for i := 0; i < r; i++ {
		m.grid[i] = make([]N, c)
	}
	return m
}

type IMat = Mat[int]

func (m Mat[N]) LUSolve(vec []N) []N {
	n := len(m.grid)
	lu := NewMat[N](n, n)
	sum := N(0)

	// LU decomp
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sum = 0
			for k := 0; k < i; k++ {
				sum += lu.grid[i][k] * lu.grid[k][j]
			}
			lu.grid[i][j] = m.grid[i][j] - sum
		}
		for j := i + 1; j < n; j++ {
			sum = 0
			for k := 0; k < i; k++ {
				sum += lu.grid[j][k] * lu.grid[k][i]
			}
			lu.grid[j][i] = (m.grid[j][i] - sum) / lu.grid[i][i]
		}
	}

	// Solve
	inter := make([]N, n)
	solution := make([]N, n)
	for i := 0; i < n; i++ {
		sum := N(0)
		for k := 0; k < i; k++ {
			sum += lu.grid[i][k] * inter[k]
		}
		inter[i] = vec[i] - sum
	}
	for i := n - 1; i >= 0; i-- {
		sum := N(0)
		for k := i + 1; k < n; k++ {
			sum += lu.grid[i][k] * solution[k]
		}
		solution[i] = (inter[i] - sum) / lu.grid[i][i]
	}
	return solution
}
