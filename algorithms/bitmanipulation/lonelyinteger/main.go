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
	strconv.Atoi(s.Text())
	var xor int

	s.Scan()
	for _, v := range strings.Split(s.Text(), " ") {
		temp, _ := strconv.Atoi(v)
		xor ^= temp
	}
	fmt.Println(xor)
}
