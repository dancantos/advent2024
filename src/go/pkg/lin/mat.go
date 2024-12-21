package lin

import "golang.org/x/exp/constraints"

type Mat[N constraints.Float] [][]N

func NewMat[N constraints.Float](r, c int) Mat[N] {
	m := Mat[N]{}
	m = make([][]N, r)
	for i := 0; i < r; i++ {
		m[i] = make([]N, c)
	}
	return m
}

func (m Mat[N]) LUSolve(vec []N) []N {
	n := len(m)
	lu := NewMat[N](n, n)
	sum := N(0)

	// LU decomp
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sum = 0
			for k := 0; k < i; k++ {
				sum += lu[i][k] * lu[k][j]
			}
			lu[i][j] = m[i][j] - sum
		}
		for j := i + 1; j < n; j++ {
			sum = 0
			for k := 0; k < i; k++ {
				sum += lu[j][k] * lu[k][i]
			}
			lu[j][i] = (m[j][i] - sum) / lu[i][i]
		}
	}

	// Solve
	inter := make([]N, n)
	solution := make([]N, n)
	for i := 0; i < n; i++ {
		sum := N(0)
		for k := 0; k < i; k++ {
			sum += lu[i][k] * inter[k]
		}
		inter[i] = vec[i] - sum
	}
	for i := n - 1; i >= 0; i-- {
		sum := N(0)
		for k := i + 1; k < n; k++ {
			sum += lu[i][k] * solution[k]
		}
		solution[i] = (inter[i] - sum) / lu[i][i]
	}
	return solution
}
