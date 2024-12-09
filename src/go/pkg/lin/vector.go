package lin

import "golang.org/x/exp/constraints"

type Vec[N constraints.Integer] struct{ x, y N }
