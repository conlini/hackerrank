package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	s.Scan()
	temp := s.Text()
	a := make([]int, n)
	for i, v := range strings.Split(temp, " ") {
		ai, _ := strconv.Atoi(v)
		a[i] = ai
	}
	fmt.Println(min(a))
}

func min(array []int) int {
	sum := sumof(array)
	if sum%2 == 0 {
		return 0
	} else {
		if len(array) > 1 {
			return 1
		} else {
			return -1
		}
	}

}

func sumof(a []int) int {
	var answer int
	for _, i := range a {
		answer += i
	}
	return answer
}
