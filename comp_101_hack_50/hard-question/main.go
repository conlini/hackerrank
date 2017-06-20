package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	buff := new(bytes.Buffer)
	max := int(math.Pow(10, 5))
	buff.Grow(max) // assuming each character is 1 byte..
	s.Buffer(buff.Bytes(), max)
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	s.Scan()
	vinc := s.Text()
	s.Scan()
	cat := s.Text()
	fmt.Println(getScore(n, vinc, cat))
}

func getScore(n int, vince, cath string) int {
	if vince == cath {
		return 0
	}
	var score = 0
	for i := 0; i < n; i++ {
		if string(vince[i]) != "." && vince[i] != cath[i] {
			score++
		}
	}

	return score
}
