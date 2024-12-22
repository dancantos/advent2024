package main

type sparseGrid struct {
	size      vec
	obstacles map[vec]struct{}
}

func newSparseGrid(size vec) sparseGrid {
	return sparseGrid{
		size:      size,
		obstacles: make(map[vec]struct{}),
	}
}

func (s sparseGrid) walk(st state, callback func(state)) bool {
	current := st
	var next state
	maxIter := s.size.x * s.size.y * 4
	var i int
	for i = 0; i < maxIter; i++ {
		next = current.next()
		if _, obstacle := s.obstacles[next.p]; obstacle {
			current = current.rotate()
			continue
		}
		if !inside(next.p, s.size) {
			return true
		}
		current = next
		callback(current)
	}
	return i < maxIter
}

type (
	vec   struct{ x, y int }
	state struct{ p, v vec }
)

type path struct {
	size    vec
	visited []uint64
}

func newPath(size vec) path {
	cels := size.x * size.y
	rows := cels / 64
	if cels%64 != 0 {
		rows++
	}
	bitMask := make([]uint64, rows)
	return path{
		size:    size,
		visited: bitMask,
	}
}

func (p path) visit(pos vec) {
	bitPos := uint64(pos.y*p.size.y + pos.x)
	row, entry := bitPos/64, bitPos%64
	p.visited[row] = p.visited[row] | 1<<entry
}

func (p path) hasVisited(pos vec) bool {
	bitPos := uint64(pos.y*p.size.y + pos.x)
	row, entry := bitPos/64, bitPos%64
	return p.visited[row]&(1<<entry) > 0
}

func (p path) countVisited() int {
	count := 0
	for _, i := range p.visited {
		// fmt.Println(i)
		count += countBits(i)
	}
	return count
}

// func (p path) printVisited() {
// 	head = 0
// 	for i := 0; i < p.size.y; i++ {

// 		for j := 0; i < p.size.x; j++ {
// 			bitPos := uint64(i*p.size.x + i)
// 			row, entry := bitPos/64, bitPos%64
// 		}
// 	}
// 	for _, i := range p.visited {
// 		fmt.Printf("%064b\n", i)
// 	}
// }

func countBits(b uint64) int {
	var count int
	for count = 0; b != 0; count++ {
		b &= b - 1
	}
	return count
}

func (s state) next() state {
	return state{vec{s.p.x + s.v.x, s.p.y + s.v.y}, s.v}
}

func (s state) rotate() state {
	return state{s.p, vec{-s.v.y, s.v.x}}
}

func rotate(v vec) vec {
	return vec{-v.y, v.x}
}

func backtrack(s state) state {
	return state{vec{s.p.x - s.v.x, s.p.y - s.v.y}, vec{-s.v.y, s.v.x}}
}

func inside(p, size vec) bool {
	return 0 <= p.x && 0 <= p.y && p.x < size.x && p.y < size.y
}
