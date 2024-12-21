package main

type filler struct {
	horizon map[vec]struct{}
	filled  map[vec]struct{}
}
