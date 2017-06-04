package main

import "fmt"

func main() {
	var in string
	fmt.Scanf("%s", &in)
	fmt.Println(reduce(in, 0))
}

// recursively reduce immediate 2 characters
func reduce(input string, startPos int) string {
	if len(input) == 0 {
		return "Empty String"
	} else {
		if len(input) >= startPos+2 {
			if input[startPos] == input[startPos+1] && len(input) >= startPos+2 {

				return reduce(input[0:startPos]+input[startPos+2:], 0)

			} else {
				return reduce(input, startPos+1)
			}
		} else {
			return input
		}
	}

}
