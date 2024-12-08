package it

import "iter"

func Count[V any](in iter.Seq[V]) int {
	count := 0
	for range in {
		count++
	}
	return count
}

func Filter[V any](in iter.Seq[V], predicate func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		in(func(v V) bool { return !predicate(v) || yield(v) })
	}
}

func Map[V1, V2 any](in iter.Seq[V1], mapFn func(V1) V2) iter.Seq[V2] {
	return func(yield func(V2) bool) {
		in(func(v V1) bool { return yield(mapFn(v)) })
	}
}

func Reduce[V, R any](in iter.Seq[V], init R, reducer func(R, V) R) R {
	result := init
	for v := range in {
		result = reducer(result, v)
	}
	return result
}

func SlicePairs[T any](slice []T) iter.Seq2[T, T] {
	return func(yield func(T, T) bool) {
		for i := 0; i < len(slice); i++ {
			for j := i + 1; j < len(slice); j++ {
				if !yield(slice[i], slice[j]) {
					return
				}
			}
		}
	}
}
