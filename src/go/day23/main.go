package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"iter"
	"maps"
	"slices"
	"strings"

	"github.com/dancantos/advent2024/src/go/pkg/it"
)

//go:embed input.txt
var input []byte

var data = readInput(bytes.NewReader(input))

func readInput(r io.Reader) graph {
	g := make(graph)
	for line := range it.ReadLines(r) {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		split := strings.Split(line, "-")
		g.addEdge(split[0], split[1])
	}
	return g
}

func main() {
	fmt.Printf("Puzzle 1: complete 3 graphs with 't': %d\n", puzzle1(data))
	fmt.Printf("Puzzle 2: maximal subgraph password: %s\n", puzzle2(data))
}

func puzzle1(g graph) int {
	count := 0
	for complete3Graph := range findNCompleteSubgraphs(g, 3) {
		if complete3Graph[0][0] == 't' || complete3Graph[1][0] == 't' || complete3Graph[2][0] == 't' {
			count++
		}
	}
	return count
}

func puzzle2(g graph) string {
	g = findMaximalSubgraph(g)
	sorted := slices.Sorted(maps.Keys(g))
	return strings.Join(sorted, ",")
}

func findMaximalSubgraph(g graph) graph {
	size := 3
	nextGraph := make(graph)
	for {
		count := 0
		for subGraph := range findNCompleteSubgraphs(g, size) {
			count++
			for _, node := range subGraph {
				nextGraph[node] = g[node]
			}
		}
		if count == 0 {
			return g
		}

		g, nextGraph = nextGraph, make(graph)
		size++
	}
}

func findNCompleteSubgraphs(g graph, n int) iter.Seq[[]string] {
	visited := make(set[string])
	return func(yield func([]string) bool) {
		for node, targets := range g {
			visited.add(node)
			for tuple := range targets.tuples(n - 1) {
				if visited.hasAny(tuple...) {
					continue
				}
				nodeWithTuple := append(tuple, node)
				if g.complete(nodeWithTuple...) && !yield(nodeWithTuple) {
					return
				}
			}
		}
	}
}

func find3CompleteSubgraphs(g graph) iter.Seq[[3]string] {
	visited := make(set[string])
	return func(yield func([3]string) bool) {
		for node, targets := range g {
			visited.add(node)
			for pair := range targets.pairs() {
				if visited.has(pair.One) || visited.has(pair.Two) {
					continue
				}
				if g.complete(node, pair.One, pair.Two) && !yield([3]string{node, pair.One, pair.Two}) {
					return
				}
			}
		}
	}
}

type (
	set[T comparable] map[T]any
	graph             map[string]set[string]
)

func (s set[T]) add(t T) {
	s[t] = struct{}{}
}

func (s set[T]) has(t T) bool {
	_, has := s[t]
	return has
}

func (s set[T]) hasAny(items ...T) bool {
	for _, t := range items {
		if _, has := s[t]; has {
			return true
		}
	}
	return false
}

func (s set[T]) pairs() iter.Seq[pair[T]] {
	return func(yield func(pair[T]) bool) {
		items := slices.Collect(maps.Keys(s))
		for i := 0; i < len(items); i++ {
			for j := i + 1; j < len(items); j++ {
				if !yield(pair[T]{items[i], items[j]}) {
					return
				}
			}
		}
	}
}

func (s set[T]) tuples(k int) iter.Seq[[]T] {
	tmp := make([]T, 0, k)
	return _tuples(slices.Collect(maps.Keys(s)), tmp, k, 0)
}

func _tuples[T any](items []T, tmp []T, k, left int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		if k == 0 {
			if !yield(tmp) {
				return
			}
		}
		for i := left; i < len(items); i++ {
			tmp = append(tmp, items[i])
			for t := range _tuples(items, tmp, k-1, i+1) {
				if !yield(t) {
					return
				}
			}
			tmp = tmp[:len(tmp)-1] // pop
		}
	}
}

func (g graph) addEdge(from, to string) {
	// add edge from-to
	targets, exists := g[from]
	if !exists {
		targets = make(set[string], 0)
		g[from] = targets
	}
	targets.add(to)

	// add edge to-from
	targets, exists = g[to]
	if !exists {
		targets = make(set[string], 0)
		g[to] = targets
	}
	targets.add(from)
}

func (g graph) complete(nodes ...string) bool {
	for i := 0; i < len(nodes); i++ {
		targetSet := g[nodes[i]]
		for j := i + 1; j < len(nodes); j++ {
			// check each target node exists in the target set
			if _, exists := targetSet[nodes[j]]; !exists {
				return false
			}
		}
	}
	return true
}

type pair[T any] struct {
	One, Two T
}
