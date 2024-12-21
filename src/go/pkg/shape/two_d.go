package shape

import (
	"iter"

	"github.com/dancantos/advent2024/src/go/pkg/lin"
)

type Box struct {
	MinX, MinY, MaxX, MaxY int
}

func (b Box) Iter() iter.Seq[lin.IVec] {
	return func(yield func(lin.IVec) bool) {
		for x := b.MinX; x <= b.MaxX; x++ {
			for y := b.MinY; y <= b.MaxY; y++ {
				if !yield(lin.IVec{x, y}) {
					return
				}
			}
		}
	}
}
