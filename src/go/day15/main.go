package main

import (
	_ "embed"
	"fmt"
	"io"
	"os"
	"strings"
)

//go:embed input.txt
var input string

var c, robotPos, commands = readInput(input)

func readInput(in string) (chart, vec, []int) {
	split := strings.Split(in, "\n\n")
	c, start := readChart(split[0])
	return c, start, readCommands(split[1])
}

func readChart(in string) (chart, vec) {
	lines := strings.Split(in, "\n")
	result := make(chart, len(lines))
	var start vec
	for i, line := range lines {
		result[i] = make([]rune, len(line))
		for j, char := range line {
			result[i][j] = char
			if char == ROBOT {
				start.x, start.y = j, i
			}
		}
	}
	return result, start
}

func readCommands(in string) []int {
	result := make([]int, 0, len(in))
	for _, char := range in {
		switch char {
		case '<':
			result = append(result, LEFT)
		case '^':
			result = append(result, UP)
		case 'v':
			result = append(result, DOWN)
		case '>':
			result = append(result, RIGHT)
		}
	}
	return result
}

func main() {
	const sampleInput = `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`

	c, _, cmds := readInput(sampleInput)
	c, startPos := c.scaleUp()
	fmt.Println(startPos)
	c.print(os.Stdout)
	fmt.Println(cmds)
}

func puzzle1(c chart, startPos vec, cmds []int) int {
	for _, cmd := range commands {
		robotPos, _ = c.moveRobot(startPos, cmd)
	}
	c.print(os.Stdout)
	return c.sumCoords()
}

type (
	chart [][]rune
	vec   struct{ x, y int }
)

func (c chart) print(out io.Writer) {
	for _, row := range c {
		fmt.Fprintln(out, string(row))
	}
}

func (c chart) sumCoords() int {
	sum := 0
	for y, row := range c {
		for x, char := range row {
			if char == BOX || char == BOX_LEFT {
				sum += 100*y + x
			}
		}
	}
	return sum
}

func (c chart) get(x, y int) rune {
	return c[y][x]
}

func (c chart) swap(p1, p2 vec) {
	c[p1.y][p1.x], c[p2.y][p2.x] = c[p2.y][p2.x], c[p1.y][p1.x]
}

func (c chart) scaleUp() (chart, vec) {
	result := make(chart, len(c))
	var newStart vec
	for y, row := range c {
		newRow := make([]rune, 2*len(row))
		for x, char := range row {
			switch char {
			case BOX:
				newRow[2*x] = BOX_LEFT
				newRow[2*x+1] = BOX_RIGHT
			case ROBOT:
				newRow[2*x] = ROBOT
				newRow[2*x+1] = EMPTY
				newStart = vec{2 * x, y}
			case EMPTY, WALL:
				newRow[2*x] = char
				newRow[2*x+1] = char
			}
		}
		result[y] = newRow
	}
	return result, newStart
}

func (c chart) moveRobot(robot vec, command int) (vec, bool) {
	nextPos := robot
	switch command {
	case UP:
		nextPos.y--
	case DOWN:
		nextPos.y++
	case LEFT:
		nextPos.x--
	case RIGHT:
		nextPos.x++
	}

	switch c.get(nextPos.x, nextPos.y) {
	case WALL:
		return robot, false
	case BOX:
		_, moved := c.moveRobot(nextPos, command)
		if moved {
			c.swap(robot, nextPos)
			return nextPos, true
		}
		return robot, false
	case BOX_LEFT:
		// lookahead

		// attempt to move the left and right
		_, moveLeft := c.moveRobot(nextPos, command)
		rightPos := vec{nextPos.x + 1, nextPos.y}
		_, moveRight := c.moveRobot(rightPos, command)
		if moveLeft && moveRight {
			c.swap(robot, nextPos)
			return nextPos, true
		}
		return robot, false
	case BOX_RIGHT:
		// attempt to move the left and right
		_, moveRight := c.moveRobot(nextPos, command)
		leftPos := vec{nextPos.x + 1, nextPos.y}
		_, moveLeft := c.moveRobot(leftPos, command)
		if moveLeft && moveRight {
			c.swap(robot, nextPos)
			return nextPos, true
		}
		return robot, false
	default: // EMPTY
		c.swap(robot, nextPos)
		return nextPos, true
	}
}

// robot move commands
const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

// chart items
const (
	EMPTY     = '.'
	WALL      = '#'
	BOX       = 'O'
	BOX_LEFT  = '['
	BOX_RIGHT = ']'
	ROBOT     = '@'
)
