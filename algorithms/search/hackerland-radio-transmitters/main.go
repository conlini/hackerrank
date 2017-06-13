package main

func main() {}

func search(x []int, k int) int {
	i := 0
	basePos := 0
	for i < len(x)-1 {
		if x[i]-x[basePos] < k {

		} else if x[i]-x[basePos] == k {

		} else {
		}
	}
	return -1
}
