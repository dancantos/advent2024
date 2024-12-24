package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/dancantos/advent2024/src/go/pkg/must"
)

func main() {
	eval(equations, variables)
	fmt.Printf("Puzzle1: eval gates: %d\n", computeResult(variables))

	result, iss := detectAllBitAdder(equations)
	if iss != "" {
		fmt.Println(iss)
	}
	slices.Sort(result)
	fmt.Printf("Puzzle2: swaps sorted: %s\n", strings.Join(result, ","))
}

type equation struct {
	v1, v2 string
	out    string
	op     int
}

type issue string

func issuef(s string, args ...any) issue {
	return issue(fmt.Sprintf(s, args...))
}

func detectAllBitAdder(eqs []equation) ([]string, issue) {
	swaps := []string{}
	bit := 0
	eqMap, _, maxBit := toEquationMap(eqs)

	// swaps are added manually by inspecting issue output
	swaps = append(swaps, eqMap.swap("z06", "fkp")...)
	swaps = append(swaps, eqMap.swap("z11", "ngr")...)
	swaps = append(swaps, eqMap.swap("z31", "mfm")...)
	swaps = append(swaps, eqMap.swap("bpt", "krj")...)

	// fmt.Println(eqMap, maxBit)

	// first half adder for bit 0
	bitStr := paddedString(bit)
	xstr, ystr, zExpected := "x"+bitStr, "y"+bitStr, "z"+bitStr
	iss, result, carry := detectHalfAdder(eqMap, xstr, ystr)
	if iss != "" {
		return nil, iss
	}
	if result != zExpected {
		return nil, issuef("adder output is not expected (expected=%s, got=%s)", zExpected, result)
	}

	// last result is just a carry, check it later
	for bit < maxBit-1 {
		bit++
		bitStr = paddedString(bit)

		// full adder detection
		xstr, ystr, zExpected = "x"+bitStr, "y"+bitStr, "z"+bitStr
		iss, result, carry = detectFullAdder(eqMap, xstr, ystr, carry)
		if iss != "" {
			return nil, iss
		}
		if result != zExpected {
			return nil, issuef("(%v,%v) adder output is not expected (expected=%s, got=%s)", zExpected, result)
		}
	}

	bit++
	bitStr = paddedString(bit)
	zExpected = "z" + bitStr
	if carry != zExpected {
		return nil, issuef("(%v,%v) adder final carry is not expected (expected=%s, got=%s)", zExpected, result)
	}

	return swaps, ""
}

func paddedString(bit int) string {
	bitStr := strconv.Itoa(bit)
	if len(bitStr) == 1 {
		bitStr = "0" + bitStr
	}
	return bitStr
}

func detectFullAdder(m equationMap, v1, v2, carry string) (issue, string, string) {
	iss, iresult, icarry := detectHalfAdder(m, v1, v2)
	if iss != "" {
		return issuef("(%s,%s): first half adder malformed: %s", v1, v2, iss), "", ""
	}
	iss, result, nextCarry := detectHalfAdder(m, iresult, carry)
	if iss != "" {
		return issuef("(%s,%s): second half adder malformed: %s", v1, v2, iss), "", ""
	}
	orEqs, hasOr := m.get(icarry, nextCarry)
	if !hasOr || len(orEqs) != 1 || orEqs[0].op != OR {
		return issuef("(%s,%s): failed to find carry OR op (expected %s OR %s)", v1, v2, icarry, nextCarry), "", ""
	}
	return "", result, orEqs[0].out
}

func detectHalfAdder(m equationMap, v1, v2 string) (issue, string, string) {
	equations, ok := m.get(v1, v2)
	if !ok {
		return issuef("equations not found for inputs %s,%s", v1, v2), "", ""
	}
	if len(equations) != 2 {
		return issuef("incorrect equation count for inputs %s,%s", v1, v2), "", ""
	}

	var andEq, xorEq equation
	switch {
	case equations[0].op == AND && equations[1].op == XOR:
		andEq, xorEq = equations[0], equations[1]
	case equations[0].op == XOR && equations[1].op == AND:
		xorEq, andEq = equations[0], equations[1]
	default:
		return issue("incorrect equation operations for inputs " + v1 + "," + v2), "", ""
	}

	return "", xorEq.out, andEq.out
}

type (
	inputPair   struct{ v1, v2 string }
	equationMap map[inputPair][]equation
)

func toEquationMap(equations []equation) (equationMap, map[string]equation, int) {
	result := make(equationMap)
	outputMap := make(map[string]equation)
	max := 0
	for _, eq := range equations {
		v1, v2 := eq.v1, eq.v2
		if v2 < v1 {
			v1, v2 = v2, v1
		}
		result[inputPair{v1, v2}] = append(result[inputPair{v1, v2}], eq)
		outputMap[eq.out] = eq
		if eq.out[0] == 'z' {
			candidateMax, err := strconv.Atoi(eq.out[1:])
			if err == nil && candidateMax > max {
				max = candidateMax
			}
		}
	}
	return result, outputMap, max
}

func (em equationMap) get(v1, v2 string) ([]equation, bool) {
	if v2 < v1 {
		v1, v2 = v2, v1
	}
	// fmt.Println(v1, v2)
	eqs, exists := em[inputPair{v1, v2}]
	return eqs, exists
}

func (em equationMap) set(v1, v2 string, i int, eq equation) {
	if v2 < v1 {
		v1, v2 = v2, v1
	}
	// fmt.Println(v1, v2)
	em[inputPair{v1, v2}][i] = eq
}

func (em equationMap) swap(out1, out2 string) []string {
	var eq1, eq2 equation
	i1, i2 := -1, -1
outer:
	for _, eqList := range em {
		for i, eq := range eqList {
			if eq.out == out1 {
				eq1 = eq
				i1 = i
			} else if eq.out == out2 {
				eq2 = eq
				i2 = i
			}
			if i1 >= 0 && i2 >= 0 {
				break outer
			}
		}
	}

	eq1.out, eq2.out = eq2.out, eq1.out
	// overwrite map entries
	em.set(eq1.v1, eq1.v2, i1, eq1)
	em.set(eq2.v1, eq2.v2, i2, eq2)

	return []string{out1, out2}
}

func eval(equations []equation, variables map[string]int) {
	evaluated := make(map[int]struct{})
	current := -1
	for len(evaluated) < len(equations) {
		// scan until we find an unevalled equation
		var eq equation
		for {
			current++
			if current == len(equations) {
				current = 0
			}
			if _, evalled := evaluated[current]; !evalled {
				break
			}
		}

		// attempt to eval in
		eq = equations[current]
		var1, var1Exists := variables[eq.v1]
		var2, var2Exists := variables[eq.v2]
		if !var1Exists || !var2Exists {
			continue
		}

		var result int
		switch eq.op {
		case AND:
			result = var1 & var2
		case OR:
			result = var1 | var2
		case XOR:
			result = var1 ^ var2
		}
		variables[eq.out] = result
		evaluated[current] = struct{}{}
	}
}

func computeResult(variables map[string]int) int {
	var result int
	for k, v := range variables {
		if k[0] == 'z' {
			bitPos, err := strconv.Atoi(k[1:3])
			if err != nil {
				continue
			}
			result |= v << bitPos
		}
	}
	return result
}

//go:embed input.txt
var input string

var variables, equations = readInput(input)

var (
	variableRegex = regexp.MustCompile(`([a-z][0-9]{2}): (0|1)`)
	equationRegex = regexp.MustCompile(`([a-z0-9]{3}) (AND|OR|XOR) ([a-z0-9]{3}) -> ([a-z0-9]{3})`)
)

func readInput(in string) (map[string]int, []equation) {
	split := strings.Split(in, "\n\n") // split on the double newline
	return readVariables(split[0]), readEquations(split[1])
}

func readVariables(in string) map[string]int {
	result := make(map[string]int)
	for _, match := range variableRegex.FindAllStringSubmatch(in, -1) {
		result[match[1]] = must.Return(strconv.Atoi(match[2]))
	}
	return result
}

func readEquations(in string) []equation {
	matches := equationRegex.FindAllStringSubmatch(in, -1)
	result := make([]equation, len(matches))
	for i, match := range matches {
		result[i] = equation{
			v1:  match[1],
			v2:  match[3],
			op:  opMap[match[2]],
			out: match[4],
		}
	}
	return result
}

// operations
const (
	AND = iota
	OR
	XOR
)

var opMap = map[string]int{
	"AND": AND,
	"OR":  OR,
	"XOR": XOR,
}
