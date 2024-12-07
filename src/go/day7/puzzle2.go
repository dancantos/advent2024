package main

import (
	"math"
)

func isSolveableWithConcat(eq equation) bool {
	if eq.target < 0 {
		return false
	}
	if len(eq.terms) == 1 {
		return eq.target == eq.terms[0]
	}
	last := eq.terms[len(eq.terms)-1]
	// divide target by last
	if eq.target%last == 0 && isSolveableWithConcat(equation{eq.target / last, eq.terms[:len(eq.terms)-1]}) {
		return true
	}

	// check unconcat
	// x1, x2 := unconcat(eq.target, last), _unconcat(eq.target, last)
	// if x1 != x2 {
	// 	panic(fmt.Errorf("unconcat(%d, %d) = %d, expected %d", eq.target, last, x1, x2))
	// }
	if unconcat := unconcat(eq.target, last); unconcat < eq.target && isSolveableWithConcat(equation{unconcat, eq.terms[:len(eq.terms)-1]}) {
		return true
	}

	// subtract last target
	if isSolveableWithConcat(equation{eq.target - last, eq.terms[:len(eq.terms)-1]}) {
		return true
	}
	return false
}

func unconcat(a, b int) int {
	if a == b {
		return 0
	}
	order := math.Ceil(math.Log10(float64(b) + 0.1))
	exp := int(math.Pow10(int(order)))
	if a%exp == b {
		return (a - b) / exp
	}
	return a
}
