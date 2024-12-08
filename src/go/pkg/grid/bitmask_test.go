package grid_test

import (
	"testing"

	"github.com/dancantos/advent2024/src/go/pkg/grid"
)

func TestBitmaskSet(t *testing.T) {
	g := grid.NewBitmask(10, 10)
	g.Set(1, 1)
	g.Set(2, 2)
	g.Set(2, 3)
	g.Set(9, 8) // hit the second row

	if !g.IsSet(1, 1) || !g.IsSet(2, 2) || !g.IsSet(2, 3) || !g.IsSet(9, 8) {
		t.Error("IsSet failed to set bits")
	}

	if g.CountSet() != 4 {
		t.Error("IsSet is overloading bits")
	}
}

func TestBitmaskUnset(t *testing.T) {
	g := grid.NewBitmask(10, 10)
	g.Set(1, 1)
	g.Set(2, 2)
	g.Set(2, 3)
	g.Set(9, 7)
	g.Set(9, 8)

	g.Unset(2, 2)
	g.Unset(9, 7)
	g.Unset(9, 6) // unset an already unset entry should do nothing

	if !g.IsSet(1, 1) || !g.IsSet(2, 3) || !g.IsSet(9, 8) {
		t.Error("Unset deleted incorrect bits")
	}

	if g.IsSet(2, 2) || g.IsSet(9, 7) {
		t.Error("Unset failed to delete the target bits")
	}

	if g.CountSet() != 3 {
		t.Error("Unset fails to decrement the total bit count")
	}
}

func TestBitmaskFlip(t *testing.T) {
	g := grid.NewBitmask(10, 10)
	g.Set(1, 1)
	g.Set(2, 2)
	g.Set(9, 7)
	g.Set(9, 8)

	g.Flip(2, 2)
	g.Flip(2, 3)
	g.Flip(9, 7)
	g.Flip(9, 9)

	if !g.IsSet(1, 1) || !g.IsSet(2, 3) || !g.IsSet(9, 8) || !g.IsSet(9, 9) {
		t.Error("Flip incorrectly unsets bits")
	}

	if g.IsSet(2, 2) || g.IsSet(9, 7) {
		t.Error("Flip incorrectly sets bits")
	}

	if g.CountSet() != 4 {
		t.Error("Unset fails to decrement the total bit count")
	}
}
