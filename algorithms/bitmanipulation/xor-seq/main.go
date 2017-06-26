package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func G(n int64) int64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	div := n / 4
	if n%4 == 0 {

		if div%2 == 0 {
			return n
		} else {
			return n + 2
		}
	} else {
		if div%2 == 0 {
			switch n % 4 {
			case 1:
				return n
			case 2, 3:
				return 2
			}
		} else {
			switch n % 4 {
			case 1:
				return n + 2
			case 2, 3:
				return 0
			}
		}
	}
	return 0
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	q, _ := strconv.Atoi(s.Text())
	A := make([]int64, 0)
	sol := make([]int64, 0)
	A = append(A, 0)
	sol = append(sol, 0)
	for i := 0; i < q; i++ {
		s.Scan()
		temp := strings.Split(s.Text(), " ")
		L, _ := strconv.ParseInt(temp[0], 10, 64)
		R, _ := strconv.ParseInt(temp[1], 10, 64)
		//R = R + 1
		fmt.Println(G(R) ^ G(L-1))
	}
}
