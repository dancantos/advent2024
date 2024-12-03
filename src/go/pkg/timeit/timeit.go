package timeit

import (
	"fmt"
	"math"
	"slices"
	"time"

	"github.com/dancantos/advent2024/src/go/pkg/it"
)

// Run yields to a for loop N times and prints time stats when completed.
func Run(count int) func(func() bool) {
	return func(yield func() bool) {
		deltas := make([]time.Duration, count)
		defer stats(deltas)
		start := time.Now()
		for i := 0; i < count; i++ {
			if !yield() {
				break
			}
			deltas[i] = time.Since(start)
		}
	}
}

func stats(deltas []time.Duration) {
	// fmt.Println(deltas)
	min := it.Reduce(slices.Values(deltas), time.Duration(math.MaxInt64), less)
	max := it.Reduce(slices.Values(deltas), 0, greater)
	sum := it.Reduce(slices.Values(deltas), 0, sum)
	mean := sum / time.Duration(len(deltas))

	fmt.Println("=== Performance Stats ===")
	fmt.Printf(" min:  %s\n", min)
	fmt.Printf(" max:  %s\n", max)
	fmt.Printf(" mean: %s\n", mean)
}

func less(a, b time.Duration) time.Duration {
	if a < b {
		return a
	}
	return b
}

func greater(a, b time.Duration) time.Duration {
	if a > b {
		return a
	}
	return b
}

func sum(a, b time.Duration) time.Duration {
	return a + b
}
