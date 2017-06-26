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
	temp := strings.Split(s.Text(), " ")
	n, _ := strconv.ParseUint(temp[0], 10, 64)
	m, _ := strconv.ParseUint(temp[1], 10, 64)
	k, _ := strconv.Atoi(temp[2])
	fmt.Println(count)
}
