package main

import "fmt"

func main() {
	s := "abcd"
	recurse(s)
}

func recurse(s string) {
	occupied := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		//		fmt.Println("occupied", occupied)
		out := string(s[i])
		occupied[i] = true
		permute(out, 0, occupied, s)
		occupied[i] = false
	}
}

func permute(out string, position int, occupied []bool, s string) {
	//	fmt.Println(out, position, occupied, s)
	if len(out) == len(s) {
		fmt.Println("Permutation", out)
		return
	}
	for i := 0; i < len(s); i++ {
		if !occupied[i] {
			occupied[i] = true
			out = out + string(s[i])
			permute(out, position+1, occupied, s)
			occupied[i] = false
			n := len(out) - 1
			out = out[:n]
		}
	}
}
