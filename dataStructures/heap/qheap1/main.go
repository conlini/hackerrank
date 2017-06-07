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
	items := make([]int, 0)
	prints := make([]int, 0)
	for i := 0; i < n; i++ {
		s.Scan()
		vals := strings.Split(s.Text(), " ")
		switch vals[0] {
		case "1":
			v, _ := strconv.Atoi(vals[1])
			items = append(items, v)
			items = sort(items)
		case "2":
			v, _ := strconv.Atoi(vals[1])
			items = search_remove(items, v)
		case "3":
			prints = append(prints, items[0])
		}
	}
	for _, i := range prints {
		fmt.Println(i)
	}

}

func search_remove(input []int, v int) []int {
	in := input[:]
	mid := len(in) / 2
	//fmt.Println(mid, in)
	for in[mid] != v {
		//fmt.Println(in[mid])
		if in[mid] < v {
			in = in[mid+1:]
			//fmt.Println(in)

		} else {
			in = in[0:mid]
		}
		//fmt.Println(mid, in)
		mid = len(in) / 2
	}
	for i := mid; i < len(input)-1; i++ {
		input[i] = input[i+1]
	}
	return input[:len(input)-1]
}

func sort(in []int) []int {
	//	//fmt.Println("sorting", in)
	if len(in) > 1 {
		mid := len(in) / 2
		l := sort(in[0:mid])
		r := sort(in[mid:])
		return merge(l, r)

	} else {
		return in
	}

}

func merge(l, r []int) []int {
	//	//fmt.Println("merging", l, r)
	if len(l) == 0 {
		//		//fmt.Println("merged", r)
		return r
	} else if len(r) == 0 {
		//		//fmt.Println("merged", l)
		return l
	}
	total := len(l) + len(r)
	var answer = make([]int, total)
	x := 0
	y := 0
	for i := 0; i < total; i++ {
		if x < len(l) && y < len(r) {
			if l[x] < r[y] {
				answer[i] = l[x]
				x++
			} else {
				answer[i] = r[y]
				y++
			}
		} else {
			if x < len(l) {
				answer[i] = l[x]
				x++
			}
			if y < len(r) {
				answer[i] = r[y]
				y++
			}
		}
		//		//fmt.Printf("progressing, answer=%v, i=%v, leftIndex=%v, rightIndex=%v\n", answer, i, x, y)
	}
	//	//fmt.Println("merged", answer)
	return answer
}
