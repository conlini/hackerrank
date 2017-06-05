package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		n   int
		arr []int
	)

	fmt.Scanf("%d", &n)
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	temp := s.Text()
	for _, x := range strings.Split(temp, " ") {
		num, _ := strconv.Atoi(x)
		arr = append(arr, num)
	}
	var shifts int
	for j := 1; j < n; j++ {

		last := arr[j]
		//		fmt.Println("last", last)
		var sorted = false
		for i := j - 1; i >= 0; i-- {
			//			fmt.Println("comparing ", arr[i])
			if arr[i] > last {
				// num greater than last. shift right
				arr[i+1] = arr[i]
				shifts++
				//	print(arr)
			} else {
				// num is now higher than current cell. Stick it in right cell
				arr[i+1] = last
				
				sorted = true
				break
			}
		}
		if !sorted {
			arr[0] = last
		}
//		print(arr)
	}
	fmt.Println(shifts)
}

func print(arr []int) {
	b := new(bytes.Buffer)
	for i, e := range arr {
		s := strconv.Itoa(e)
		if i != 0 {
			b.Write([]byte(" "))
		}
		b.Write([]byte(s))
	}
	fmt.Println(string(b.Bytes()))
}
