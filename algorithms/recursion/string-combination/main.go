package main

import "fmt"

func main() {
	s := "abc"
	recurse(s)

}

func recurse(s string) {
	for L := 1; L <= len(s); L++ {
		occupied := make([]bool, len(s))
		for i := 0; i < len(s); i++ {
			out := string(s[i])
			occupied[i] = true
			permute(out, 0, occupied, L, s)
			occupied[i] = false

		}
	}
}

func permute(out string, position int, occupied []bool, l int, s string) {
	//	fmt.Println(out, position, occupied, s)
	if len(out) == l {
		fmt.Println("Permutation", out)
		return
	}
	for i := 0; i < len(s); i++ {
		if !occupied[i] {
			occupied[i] = true
			out = out + string(s[i])
			permute(out, position+1, occupied, l, s)
			occupied[i] = false
			n := len(out) - 1
			out = out[:n]
		}
	}
}
