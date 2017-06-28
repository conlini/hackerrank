package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var count int

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	d, _ := strconv.Atoi(s.Text())
	print2(d, d, "")
	fmt.Println(count)
}

func print(l, r int, out string) {
	if r == 0 {
		count++
		fmt.Println(out)
		return
	}
	if l == r {
		print(l-1, r, out+"(")
	} else if l == 0 {
		print(l, r-1, out+")")
	} else {
		print(l-1, r, out+"(")
		print(l, r-1, out+")")
	}

}

func print2(l, r int, out string) {
	if r == 0 {
		count++
		fmt.Println(out)
		return
	}
	if l > r {
		print(l, r-1, out+")")

	}
	if l > 0 {
		print(l-1, r, out+"(")
	}

}
