package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		m       [36]int
		rows    = 6
		columns = 6
	)
	s := bufio.NewScanner(os.Stdin)
	for i := 0; i < 6; i++ {
		s.Scan()
		temp := strings.Split(s.Text(), " ")
		for y, v := range temp {
			n, _ := strconv.Atoi(v)
			m[i*6+y] = n
		}
	}
	//	fmt.Println(m)
	var maxSum = 6 * -9
	for r := 0; r < rows-2; r++ {
		for c := 0; c < columns-2; c++ {
			i := c + (r * rows)
			indices := []int{i, i + 1, i + 2, i + 7, i + 12, i + 13, i + 14}
			var sum int
			for _, x := range indices {
				sum += m[x]
			}
			fmt.Println(sum)
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	fmt.Println(maxSum)
}
