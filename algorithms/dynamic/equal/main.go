package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	t, _ := strconv.Atoi(s.Text())
	for i := 0; i < t; i++ {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())
		chocs := make([]int, n)
		s.Scan()
		for i, v := range strings.Split(s.Text(), " ") {
			temp, _ := strconv.Atoi(v)
			chocs[i] = temp
		}
		fmt.Println(mincount(chocs))
	}
}

func mincount(n []int) int {
	// sort
	sort.Ints(n)
	// do a diff
	diff := make([]int, len(n)-1)

	for i := 1; i < len(n); i++ {
		//		fmt.Println(i)
		differ := int(math.Abs(float64(n[i] - n[i-1])))
		diff[i-1] = differ

	}
	var min int

	for i := 0; i < len(diff)-1; i++ {
		//	fmt.Println(diff, i)
		if diff[i] > 0 {
			//	fmt.Println(diff[i+1])
			diff[i+1] = diff[i+1] + diff[i]
			//	fmt.Println("after", diff)
			min += getMin(diff[i])
			//			fmt.Println("min", min)
			diff[i] = 0
		}
	}
	//	fmt.Println(diff)
	last := diff[len(diff)-1]
	//	fmt.Println("last", last)
	min += getMin(last)
	return min
}

func getMin(n int) int {
	var min int
	if n >= 5 {
		if n%5 == 0 {
			//			fmt.Println(n / 5)
			return n / 5
		} else {
			min += n / 5
			n = n % 5
		}
	}
	if n%2 == 0 {
		min += n / 2
		return min
	} else {
		min += n / 2
		min += 1
	}
	return min
}
