package timeit

import (
	"fmt"
	"math"
	"slices"
	"time"

	"github.com/dancantos/advent2024/src/go/pkg/it"
)

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

	// precision := int(math.Min(3*math.Floor(math.Log(min)/math.Log(10)/3), 9))
	// unit := []string{"ns", "µs", "ms", "s"}[precision/3]
	// div := math.Pow(10, float64(precision))

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

// def run(n: int) -> iter:
//     times = [perf_counter_ns()]
//     for i in range(n):
//         yield i
//         times.append(perf_counter_ns())
//     deltas = [b - a for a, b in zip(times[:-1], times[1:])]
//     precision = min(3 * math.floor(math.log(min(deltas)) / math.log(10) / 3), 9)
//     unit = ["ns", "μs", "ms", "s"][precision//3]
//     print("=== Performance Stats ===")
//     print(f" min:  {min(deltas)/10**precision:-10.3f}{unit}")
//     print(f" max:  {max(deltas)/10**precision:-10.3f}{unit}")
//     print(f" mean: {sum(deltas) / len(deltas)/10**precision:-10.3f}{unit}")
