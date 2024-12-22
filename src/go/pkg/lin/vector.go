package lin

import "golang.org/x/exp/constraints"

type Vec[N constraints.Integer] struct{ X, Y N }

type IVec = Vec[int]
