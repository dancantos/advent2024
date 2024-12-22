package main

import "strings"

var cmdChars = map[int]rune{
	PRESS: 'A',
	UP:    '^',
	DOWN:  'v',
	LEFT:  '<',
	RIGHT: '>',
}

func cmdString(commands []int) string {
	builder := strings.Builder{}
	for _, cmd := range commands {
		builder.WriteRune(cmdChars[cmd])
	}
	return builder.String()
}

func computeCommands(code string, indirections int) []int {
	start := terminal['A']

	// 1 robot conrtolling the terminal
	commands := moveToAndPressTerminal(start, terminal[rune(code[0])])
	for i := 0; i < len(code)-1; i++ {
		from, to := terminal[rune(code[i])], terminal[rune(code[i+1])]
		commands = moveToAndPressTerminal(from, to)
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
	commands := moveToAndPressPad(start, keypad[cmds[0]])
	for i := 0; i < len(cmds)-1; i++ {
		from, to := keypad[cmds[i]], keypad[cmds[i+1]]
		commands = append(commands, moveToAndPressPad(from, to)...)
	}
	return commands
}
