package main

import (
	"fmt"
	"regexp"
)

func main() {
	var in string
	fmt.Scanf("%s", &in)
	r := regexp.MustCompile("[A-Z]")
	match := r.FindAllString(in, -1)
	fmt.Println(len(match) + 1)
}
