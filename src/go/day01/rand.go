package main

import "math/rand"

func randArr() []int {
	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(10)
	}
	return arr
}
