package main

import "fmt"

func main() {
	var (
		in   int
		temp string
		vals []string
	)
	fmt.Scanf("%d", &in)
	for i := 0; i < in; i++ {
		fmt.Scanf("%s", &temp)
		vals = append(vals, temp)
	}
	out := sort(vals)
	for i := 0; i < in; i++ {
		fmt.Println(out[i])
	}
}

func sort(in []string) []string {
	//	fmt.Println("sorting", in)
	if len(in) > 1 {
		mid := len(in) / 2
		l := sort(in[0:mid])
		r := sort(in[mid:])
		return merge(l, r)

	} else {
		return in
	}

}
func merge(l, r []string) []string {
	//	fmt.Println("merging", l, r)
	if len(l) == 0 {
		//		fmt.Println("merged", r)
		return r
	} else if len(r) == 0 {
		//		fmt.Println("merged", l)
		return l
	}
	total := len(l) + len(r)
	var answer = make([]string, total)
	x := 0
	y := 0
	for i := 0; i < total; i++ {
		if x < len(l) && y < len(r) {
			choice, which := compare(l[x], r[y])
			switch which {
			case "left":
				answer[i] = choice
				x++
			case "right":
				answer[i] = choice
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
		//		fmt.Printf("progressing, answer=%v, i=%v, leftIndex=%v, rightIndex=%v\n", answer, i, x, y)
	}
	//	fmt.Println("merged", answer)
	return answer
}

func compare(l, r string) (string, string) {
	//	fmt.Printf("comparing %v with %v\n", l, r)
	if len(l) < len(r) {
		return l, "left"
	}
	if len(r) < len(l) {
		return r, "right"
	}
	for i := 0; i < len(l); i++ {
		if l[i] < r[i] {
			return l, "left"
		} else if r[i] < l[i] {
			return r, "right"
		}
	}
	return l, "left"
}
