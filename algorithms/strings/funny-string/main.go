package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	T, _ := strconv.Atoi(s.Text())
	var answer = make([]string, T)
	for i := 0; i < T; i++ {
		s.Scan()
		answer[i] = isFunny(s.Text())
	}
	for _, i := range answer {
		fmt.Println(i)
	}
}

func isFunny(data string) string {
	b := []byte(data)
	var leftDiff, rightDiff int
	length := len(b)
	for i := 1; i < length; i++ {
		leftDiff = int(math.Abs(float64(b[i]) - float64(b[i-1])))
		rightDiff = int(math.Abs(float64(b[length-i]) - float64(b[length-i-1])))
		if leftDiff != rightDiff {
			return "Not Funny"
		}
	}

	return "Funny"
}
