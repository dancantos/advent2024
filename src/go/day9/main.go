package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	p := process(input)
	// fmt.Println(p)
	l := left2(p)
	// fmt.Println(l)
	fmt.Println(checksum(l))
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

func left(arr []int) []int {
	head := 0
	tail := len(arr) - 1

	for head < tail {
		head = seekEmpty(arr, head)
		tail = seekFull(arr, tail)
		if head > tail {
			break
		}
		arr = swap(arr, head, tail)
		head++
		tail--
		// fmt.Println(head, tail, arr)
	}
	return arr
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

func seekEmpty(arr []int, head int) int {
	for i := head; i < len(arr); i++ {
		if arr[i] == -1 {
			return i
		}
	}
	return -1
}

func seekFull(arr []int, head int) int {
	for i := head; i >= 0; i-- {
		if arr[i] != -1 {
			return i
		}
	}
	return -1
}

func swap(arr []int, a, b int) []int {
	arr[a], arr[b] = arr[b], arr[a]
	return arr
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
