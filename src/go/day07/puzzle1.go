package main

func isSolveable(eq equation) bool {
	if len(eq.terms) == 1 {
		return eq.target == eq.terms[0]
	}
	last := eq.terms[len(eq.terms)-1]
	// divide target by last
	if eq.target%last == 0 && isSolveable(equation{eq.target / last, eq.terms[:len(eq.terms)-1]}) {
		return true
	}

	// subtract last
	if eq.target > last && isSolveable(equation{eq.target - last, eq.terms[:len(eq.terms)-1]}) {
		return true
	}

	return false
}
