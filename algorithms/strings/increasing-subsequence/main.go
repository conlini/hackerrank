package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	occupied := make([]bool, len(n))
	for i := 0; i < len(n); i++ {
		occupied[i] = true
		subsequence(4, n, i, occupied, []int{n[i]})
	}
}

func subsequence(k int, n []int, currentPos int, occupied []bool, out []int) {
	//	fmt.Printf("current %d, occupied %v, out %s\n", currentPos, occupied, out)
	if len(out) == k {
		fmt.Println(out)
		return
	}
	for i := currentPos + 1; i < len(n); i++ {
		if !occupied[i] {
			occupied[i] = true
			subsequence(k, n, i, occupied, append(out, n[i]))
			occupied[i] = false
		}
	}
}
