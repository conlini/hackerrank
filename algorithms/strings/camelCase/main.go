package main

import (
	"fmt"
	"regexp"
)

func main() {
	// use regex to identiy how many caps are present. Num words = reg match + 1
	var in string
	fmt.Scanf("%s", &in)
	r := regexp.MustCompile("[A-Z]")
	match := r.FindAllString(in, -1)
	fmt.Println(len(match) + 1)
}
