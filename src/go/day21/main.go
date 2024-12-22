package main

import (
	"fmt"
)

func main() {
	fmt.Println("Puzzle1 (2 mid-robots):", sumComplexities(codes, nums, 2))
	fmt.Println("Puzzle2 (25 mid-robots):", sumComplexities(codes, nums, 25))
}

func sumComplexities(codes []string, nums []int, robots int) int {
	m := make(memory)
	sum := 0
	for i, code := range codes {
		sum += countCommands(code, robots, m) * nums[i]
	}
	return sum
}

func countCommands(code string, depth int, m memory) int {
	// 1 robot at the terminal
	// its not worth memoizing this since no character pair re-occurs in the problem input
	commands := moveToAndPressTerminal(terminal['A'], terminal[rune(code[0])])
	for i := 0; i < len(code)-1; i++ {
		from, to := terminal[rune(code[i])], terminal[rune(code[i+1])]
		commands = append(commands, moveToAndPressTerminal(from, to)...)
	}

	// n robots in series
	return _countCommands(commands, depth, m)
}

func _countCommands(commands []int, depth int, m memory) int {
	if depth == 0 {
		// fmt.Print(cmdString(commands))
		return len(commands)
	}
	sum := 0
	var pairSum int
	var exists bool
	commands = append([]int{PRESS}, commands...)
	for i := 0; i < len(commands)-1; i++ {
		current, next := commands[i], commands[i+1]
		if pairSum, exists = m.get(vec{current, next}, depth); !exists {
			pairSum = _countCommands(moveToAndPressPad(keypad[current], keypad[next]), depth-1, m)
			m.put(vec{current, next}, depth, pairSum)
		}
		sum += pairSum
	}
	return sum
}

type vec struct{ x, y int }

func moveToAndPressTerminal(start, destination vec) []int {
	result := make([]int, 0, 5)
	if start.y == 0 && destination.x == 0 {
		// here we need to prefer up over left
		if destination.y > start.y {
			for i := 0; i < destination.y-start.y; i++ {
				result = append(result, UP)
			}
		}
		if destination.x < start.x {
			for i := 0; i < start.x-destination.x; i++ {
				result = append(result, LEFT)
			}
		}
	} else if start.x == 0 && destination.y == 0 {
		if destination.x > start.x {
			for i := 0; i < destination.x-start.x; i++ {
				result = append(result, RIGHT)
			}
		}
		if destination.y < start.y {
			for i := 0; i < start.y-destination.y; i++ {
				result = append(result, DOWN)
			}
		}
	} else {
		if destination.x < start.x {
			for i := 0; i < start.x-destination.x; i++ {
				result = append(result, LEFT)
			}
		}
		if destination.y < start.y {
			for i := 0; i < start.y-destination.y; i++ {
				result = append(result, DOWN)
			}
		}
		if destination.y > start.y {
			for i := 0; i < destination.y-start.y; i++ {
				result = append(result, UP)
			}
		}
		if destination.x > start.x {
			for i := 0; i < destination.x-start.x; i++ {
				result = append(result, RIGHT)
			}
		}
	}
	return append(result, PRESS)
}

func moveToAndPressPad(start, destination vec) []int {
	result := make([]int, 0, 5)

	if destination.x == 0 && start.y == 1 {
		// if start -> end could contain top-left forbidden spot, prefer DOWN over LEFT
		if destination.y < start.y {
			for i := 0; i < start.y-destination.y; i++ {
				result = append(result, DOWN)
			}
		}
		if destination.x < start.x {
			for i := 0; i < start.x-destination.x; i++ {
				result = append(result, LEFT)
			}
		}
	} else if start.x == 0 && destination.y == 1 {
		// if end -> start could contain top-left forbidden spot, prefer RIGHT over UP
		if destination.x > start.x {
			for i := 0; i < destination.x-start.x; i++ {
				result = append(result, RIGHT)
			}
		}
		if destination.y > start.y {
			for i := 0; i < destination.y-start.y; i++ {
				result = append(result, UP)
			}
		}
	} else {
		// otherwise prefer LEFT < DOWN < UP < RIGHT
		// this heuristic gives optimal solution
		if destination.x < start.x {
			for i := 0; i < start.x-destination.x; i++ {
				result = append(result, LEFT)
			}
		}
		if destination.y < start.y {
			for i := 0; i < start.y-destination.y; i++ {
				result = append(result, DOWN)
			}
		}
		if destination.y > start.y {
			for i := 0; i < destination.y-start.y; i++ {
				result = append(result, UP)
			}
		}
		if destination.x > start.x {
			for i := 0; i < destination.x-start.x; i++ {
				result = append(result, RIGHT)
			}
		}
	}
	return append(result, PRESS)
}

// commands that a robot can give to the next
const (
	PRESS = iota
	UP
	DOWN
	LEFT
	RIGHT
)

// keypad command locations
var keypad = map[int]vec{
	PRESS: {2, 1},
	UP:    {1, 1},
	DOWN:  {1, 0},
	LEFT:  {0, 0},
	RIGHT: {2, 0},
}

// end terminal input locations
var terminal = map[rune]vec{
	'A': {2, 0},
	'0': {1, 0},
	'1': {0, 1},
	'2': {1, 1},
	'3': {2, 1},
	'4': {0, 2},
	'5': {1, 2},
	'6': {2, 2},
	'7': {0, 3},
	'8': {1, 3},
	'9': {2, 3},
}

type memory map[vec][]int

func (m memory) put(pair vec, depth, count int) {
	depths, exists := m[pair]
	if !exists {
		depths = make([]int, 26) // 25 is max
		m[pair] = depths
	}
	depths[depth] = count
}

func (m memory) get(pair vec, depth int) (int, bool) {
	depths, exists := m[pair]
	if !exists {
		return 0, false
	}
	count := depths[depth]
	return count, count > 0
}
