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
	max := int(math.Pow(10, 5))
	buff.Grow(max)
	s.Buffer(buff.Bytes(), max)
	s.Scan()
	temp := strings.Split(s.Text(), " ")
	L, _ := strconv.Atoi(temp[0])
	R, _ := strconv.Atoi(temp[1])
	var max int

	for i := L; i <= R; i++ {
		for j := L; j <= R; j++ {
			xor := i ^ j
			if xor > max {
				max = xor
			}
		}
	}
	fmt.Println(max)
}
