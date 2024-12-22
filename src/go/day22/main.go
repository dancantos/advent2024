package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"maps"
	"strconv"

	"github.com/dancantos/advent2024/src/go/pkg/must"
)

const (
	PRUNE_MOD  = 16777216
	NUM_PRICES = 2000
)

func main() {
	fmt.Println("Puzzle 1 (count 2000th prices): ", count2000Prices(nums))
	fmt.Println("Puzzle 2 (maximize bananas):", findMaxBananas(nums))
}

func price(n int) int {
	return n % 10
}

func findMaxBananas(nums []int) int {
	m := make(memory)
	for _, n := range nums {
		memoizeDiffs(n, m)
	}
	var maxBananas int
	for bananaCount := range maps.Values(m) {
		if bananaCount > maxBananas {
			maxBananas = bananaCount
		}
	}
	return maxBananas
}

func count2000Prices(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += produceNumbers(n, 2000)
	}
	return sum
}

func memoizeDiffs(n int, m memory) {
	localBests := make(memory)
	diffs := [4]int{}

	next := process(n)
	currentPrice, nextPrice := price(n), price(next)
	diffs[0] = nextPrice - currentPrice

	n, next = next, process(next)
	currentPrice, nextPrice = nextPrice, price(next)
	diffs[1] = nextPrice - currentPrice

	n, next = next, process(next)
	currentPrice, nextPrice = nextPrice, price(next)
	diffs[2] = nextPrice - currentPrice

	n, next = next, process(next)
	currentPrice, nextPrice = nextPrice, price(next)
	diffs[3] = nextPrice - currentPrice

	cp := [4]int{}
	copy(cp[:], diffs[:])
	localBests[diffs] = nextPrice

	for i := 5; i < NUM_PRICES; i++ {
		copy(diffs[:3], diffs[1:])
		n, next = next, process(next)
		currentPrice, nextPrice = nextPrice, price(next)
		diffs[3] = nextPrice - currentPrice
		if _, exists := localBests[diffs]; !exists {
			localBests[diffs] = nextPrice
		}
	}

	for k, v := range localBests {
		m[k] += v
	}
}

func produceNumbers(seed int, nth int) int {
	for i := 0; i < nth; i++ {
		seed = process(seed)
	}
	return seed
}

func process(num int) int {
	// mul 64, mix, prune
	num = mixAndPrune(num, num*64)
	num = mixAndPrune(num, num/32)
	num = mixAndPrune(num, num*2048)
	return num
}

func mixAndPrune(secret, mixer int) int {
	secret ^= mixer
	secret %= PRUNE_MOD
	return secret
}

//go:embed input.txt
var input []byte

var nums = readInput(bytes.NewReader(input))

func readInput(r io.Reader) []int {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	result := []int{}
	for s.Scan() {
		result = append(result, must.Return(strconv.Atoi(s.Text())))
	}
	return result
}

type memory map[[4]int]int
