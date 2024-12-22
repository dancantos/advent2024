package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/dancantos/advent2024/src/go/pkg/must"
)

//go:embed input.txt
var input []byte

var codes, nums = readInput(bytes.NewReader(input))

func main() {
	// 	const test = `029A
	// 980A
	// 179A
	// 456A
	// 379A`
	// 	codes, nums := readInput(strings.NewReader(test))
	// fmt.Println(countComplexities(codes, 2))
	fmt.Println(countComplexities2(codes, 2))
	fmt.Println(countComplexities2(codes, 25))
}

func countComplexities(codes []string, indirection int) int {
	sum := 0
	for i, code := range codes {
		// fmt.Print("code '", code, "':")
		commands := computeCommands(code, indirection)
		sum += len(commands) * nums[i]
		fmt.Println()
	}
	return sum
}

func countComplexities2(codes []string, indirection int) int {
	sum := 0
	for i, code := range codes {
		sum += countCommands(code, indirection) * nums[i]
	}
	return sum
}

func countCommands(code string, indirection int) int {
	m := make(memory)
	// 1 robot at the terminal
	commands := moveToAndPressTerminal(terminal['A'], terminal[rune(code[0])])
	for i := 0; i < len(code)-1; i++ {
		from, to := terminal[rune(code[i])], terminal[rune(code[i+1])]
		commands = append(commands, moveToAndPressTerminal(from, to)...)
	}

	// n robots in series
	result := _countCommands(commands, indirection, m)
	// fmt.Println()
	return result
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
			pairSum = _countCommands(buildNextLayer(current, next), depth-1, m)
			m.put(vec{current, next}, depth, pairSum)
		}
		sum += pairSum
	}
	return sum
}

func buildNextLayer(current, next int) []int {
	return moveToAndPress(keypad[current], keypad[next])
	// return append(moveToAndPress(keypad[PRESS], keypad[current]), moveToAndPress(keypad[current], keypad[next])...)
}

func computeCommands(code string, indirections int) []int {
	start := terminal['A']

	// 1 robot conrtolling the terminal
	commands := moveToAndPressTerminal(start, terminal[rune(code[0])])
	for i := 0; i < len(code)-1; i++ {
		from, to := terminal[rune(code[i])], terminal[rune(code[i+1])]
		commands = append(commands, moveToAndPressTerminal(from, to)...)
	}

	// 2 robots controlling each other
	for i := 0; i < indirections; i++ {
		// fmt.Print(" ", i)
		commands = computeMetaCommands(commands)
	}

	// 1 me controlling the robot
	return commands
}

func computeMetaCommands(cmds []int) []int {
	start := keypad[PRESS]
	commands := moveToAndPress(start, keypad[cmds[0]])
	for i := 0; i < len(cmds)-1; i++ {
		from, to := keypad[cmds[i]], keypad[cmds[i+1]]
		commands = append(commands, moveToAndPress(from, to)...)
	}
	return commands
}

func cmdString(commands []int) string {
	builder := strings.Builder{}
	for _, cmd := range commands {
		builder.WriteRune(cmdChars[cmd])
	}
	return builder.String()
}

var numRegex = regexp.MustCompile(`\d+`)

func readInput(r io.Reader) ([]string, []int) {
	codes := []string{}
	nums := []int{}
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		codes = append(codes, s.Text())
		n := numRegex.FindString(s.Text())
		nums = append(nums, must.Return(strconv.Atoi(n)))
	}
	return codes, nums
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

func moveToAndPress(start, destination vec) []int {
	result := make([]int, 0, 5)
	if destination.x == 0 && start.y == 1 {
		// here we need to prefer up over left
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

const (
	PRESS = iota
	UP
	DOWN
	LEFT
	RIGHT
)

var keypad = map[int]vec{
	PRESS: {2, 1},
	UP:    {1, 1},
	DOWN:  {1, 0},
	LEFT:  {0, 0},
	RIGHT: {2, 0},
}

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

var cmdChars = map[int]rune{
	PRESS: 'A',
	UP:    '^',
	DOWN:  'v',
	LEFT:  '<',
	RIGHT: '>',
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
