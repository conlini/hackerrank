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
	max := int(math.Pow(10, 7))
	buff.Grow(max)
	s.Buffer(buff.Bytes(), max)
	s.Scan()
	temp := strings.Split(s.Text(), " ")
	//fmt.Println("\n\nParsed", temp)
	//n, _ := strconv.ParseUint(temp[0], 10, 64)
	k, _ := strconv.ParseInt(temp[1], 10, 64)
	//fmt.Println("\n\n number", n, k)
	super_dig := superDigit(temp[0])
	fmt.Println(superDigit(fmt.Sprintf("%d", super_dig*k, 0)))
}

func superDigit(n string) int64 {
	if len(n) == 1 {
		a, _ := strconv.ParseInt(n, 10, 64)
		return a
	}
	sum := findSum(n)
	return superDigit(fmt.Sprintf("%d", sum))
}

func findSum(n string) int64 {
	if len(n) == 1 {
		a, _ := strconv.ParseInt(n, 10, 64)
		return a
	}
	mid := len(n) / 2
	return findSum(n[:mid]) + findSum(n[mid:])
}
