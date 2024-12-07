package main

import (
	"fmt"
	"testing"
)

func TestUnconcat(t *testing.T) {
	for _, test := range []struct {
		a, b int
		want int
	}{
		{1, 1, 0},
		{12, 2, 1},
		{123, 23, 1},
		{1234, 34, 12},
		{1234, 1235, 1234},
		{1234, 1233, 1234},
		{1234, 1234, 0},
		{229107331, 1, 22910733},
	} {
		name := fmt.Sprintf("unconcat(%d,%d)", test.a, test.b)
		t.Run(name, func(t *testing.T) {
			got := unconcat(test.a, test.b)
			if got != test.want {
				t.Errorf("unconcat(%d,%d) = %d, want %d", test.a, test.b, got, test.want)
			}
		})
	}
}
