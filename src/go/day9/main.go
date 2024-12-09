package main

import (
	_ "embed"
	"fmt"

	"github.com/dancantos/advent2024/src/go/pkg/timeit"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("Puzzle 1 (fragmented checksum): %d\n", inplaceChecksum1([]byte(input)))
	for range timeit.Run(1) {
		inplaceChecksum1([]byte(input))
	}
}

func process(input string) []int {
	active := true
	result := []int{}
	count := 0
	for _, char := range input {
		i := char - '0'
		if active {
			for j := 0; j < int(i); j++ {
				result = append(result, count)
			}
			count++
		} else {
			for j := 0; j < int(i); j++ {
				result = append(result, -1)
			}
		}
		active = !active
	}
	return result
}

func left2(arr []int) []int {
	tail := len(arr) - 1
	indicator := arr[len(arr)-1]

	for tail > 0 {
		blockSize := countBlockSize(arr, tail)
		if blockSize == -1 {
			return arr
		}
		location := seekEmptyBlock2(arr, 0, blockSize, tail)
		if location == -1 {
			tail = tail - blockSize
			continue
		}
		arr = swap2(arr, location, tail, blockSize)
		tail = seekFull(arr, tail)
		indicator--
	}
	return arr
}

func seekEmptyBlock2(arr []int, head int, size int, tail int) int {
	for i := head; i < tail-size; i++ {
		if isEmpty(arr[i : i+size]) {
			return i
		}
	}
	return -1
}

func seekEmptyBlock(arr []int, head int, size int) int {
	for i := head; i < len(arr); i++ {
		if isEmpty(arr[i : i+size]) {
			return i
		}
	}
	return -1
}

func countBlockSize(arr []int, tail int) int {
	entry := arr[tail]
	count := 1
	for i := tail - 1; i >= 0; i-- {
		if arr[i] != entry {
			return count
		}
		count++
	}
	return -1
}

func isEmpty(slice []int) bool {
	for _, n := range slice {
		if n != -1 {
			return false
		}
	}
	return true
}

func swap2(arr []int, loc int, tail int, blockSize int) []int {
	// fmt.Println(arr[loc:loc+blockSize], arr[tail-blockSize+1:tail+1], loc, tail)
	copy(arr[loc:loc+blockSize], arr[tail-blockSize+1:tail+1])
	for i := tail; i >= tail-blockSize+1; i-- {
		arr[i] = -1
	}
	return arr
}

func checksum(arr []int) int {
	sum := 0
	for i, num := range arr {
		if num == -1 {
			continue
		}
		sum += i * num
	}
	return sum
}

func checksumP(s string) int {
	sum := 0
	for i, char := range s {
		sum += i * int(char-'0')
		fmt.Println(sum)
	}
	return sum
}
