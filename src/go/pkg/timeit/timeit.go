package timeit

import "time"

func Run(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Since(start)
}
