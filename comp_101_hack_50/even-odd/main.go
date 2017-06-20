package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	buff := new(bytes.Buffer)
	max := int(math.Pow(10, 11))
	buff.Grow(max)
	s.Buffer(buff.Bytes(), max)
	s.Scan()
	q, _ := strconv.Atoi(s.Text())
	answer := make([]int, q)
	for i := 0; i < q; i++ {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())
		boxes := make([]int, n)
		s.Scan()
		for j, v := range strings.Split(s.Text(), " ") {
			w, _ := strconv.Atoi(v)
			boxes[j] = w
		}
		answer[i] = makeBeautiful(n, boxes)
	}
	for _, a := range answer {
		fmt.Println(a)
	}
}

func makeBeautiful(n int, x []int) int {
	// lets start by finding position that are off and maintain an array of invalid position

	invalid := make([]int, 0)
	for i := 0; i < n; i++ {
		mod := i % 2
		if mod == 0 {
			// even bucket
			if x[i]%2 != 0 {
				invalid = append(invalid, i)
			}
		} else {
			// odd bucket
			if x[i]%2 == 0 {
				invalid = append(invalid, i)
			}
		}
	}
	if len(invalid) == 0 {
		// all positions are valid. no swaps required
		return 0
	} else if len(invalid)%2 == 1 {
		// we have odd number of errors. This can not be corrected
		return -1
	} else {
		return len(invalid) / 2
	}

	return -1
}
