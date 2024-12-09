package main

func inplaceChecksum1(input []byte) int {
	checksum := 0
	tailIdent := len(input) / 2
	head, tail := 0, len(input)-1
	counter := 0
	headIdent, tailIdent := 0, len(input)/2
	var sum int
	for head < tail {
		// add checksum at head
		// fmt.Println("add", counter, headIdent, input[head]-'0')
		counter, sum = addChecksum(counter, int(input[head]-'0'), headIdent)
		checksum += sum
		head++
		headIdent++

		// for gaps, add checksums at tail
		gapSize := int(input[head] - '0')
		for gapSize > 0 {
			tailSize := int(input[tail] - '0')
			if tailSize <= gapSize {
				// fmt.Println("add tail", counter, tailIdent, tailSize)
				counter, sum = addChecksum(counter, tailSize, tailIdent)
				checksum += sum
				tail -= 2
				gapSize -= tailSize
				tailIdent--
			} else {
				// fmt.Println("add tail", counter, tailIdent, gapSize)
				counter, sum = addChecksum(counter, gapSize, tailIdent)
				checksum += sum
				input[tail] = byte(tailSize-gapSize) + '0'
				gapSize = 0
			}
		}
		head++
	}
	final := int(input[head] - '0')
	if tail > head {
		final += int(input[tail] - '0')
	}
	_, sum = addChecksum(counter, final, headIdent)
	checksum += sum

	return checksum
}

func addChecksum(counter, count, ident int) (int, int) {
	sum := 0
	for i := counter; i < counter+count; i++ {
		// fmt.Println(i * ident)
		sum += i * ident
	}
	return counter + count, sum
}

func rearrange1(arr []int) {
	head := 0
	tail := len(arr) - 1
	for tail > 0 {
		head = seekEmpty(arr, head)
		tail = seekFull(arr, tail)
		if head > tail {
			break
		}
		swap(arr, head, tail)
		head++
		tail--
	}
}

func swap(arr []int, head, tail int) {
	arr[head], arr[tail] = arr[tail], arr[head]
}

func seekFull(arr []int, tail int) int {
	for i := tail; i > 0; i-- {
		if arr[i] >= 0 {
			return i
		}
	}
	return -1
}

func seekEmpty(arr []int, start int) int {
	for i := start; i < len(arr); i++ {
		if arr[i] == -1 {
			return i
		}
	}
	return -1
}
