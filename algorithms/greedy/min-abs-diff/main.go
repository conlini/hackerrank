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
	//	file, _ := os.Open(os.Stdin)
	//	s := bufio.NewScanner(file)
	buf := new(bytes.Buffer)
	buf.Grow(1024 * 1024)

	s.Buffer(buf.Bytes(), 1024*1024)
	s.Scan()
	n, err := strconv.Atoi(s.Text())
	if err != nil {
		fmt.Println(err)
		return
	}
	//	fmt.Println(n)

	scanned := s.Scan()
	if !scanned {
		fmt.Println("LongFile", s.Err())
		return
	}
	temp := s.Text()
	var in = make([]int, n)
	//	fmt.Println(temp)
	for i, sv := range strings.Split(temp, " ") {
		v, _ := strconv.Atoi(sv)
		in[i] = v
	}
	fmt.Println(minAbsDiff(in))
}

func minAbsDiff(copy []int) int {
	//	copy := make([]int, len(in))

	//	for i, v := range in {
	//		copy[i] = int(math.Abs(float64(v)))
	//	}
	//	sort.Ints(copy)
	var min = math.MaxInt32

	for i := 0; i < len(copy)-2; i++ {
		diff := copy[i] - copy[i+1]
		diff = int(math.Abs(float64(diff)))
		if diff == 0 {
			fmt.Println(i, copy[i+1], copy[i])
		}
		if diff < min {
			min = diff
		}
	}
	return min
}
