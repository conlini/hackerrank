package main

import (
	"bytes"
	"fmt"
)

func main() {
	var (
		n    int
		temp string
		arr  []string
	)

	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &temp)
		arr = append(arr, temp)
	}

	b := new(bytes.Buffer)
	for i := n - 1; i >= 0; i-- {
		b.Write([]byte(arr[i]))
		b.Write([]byte(" "))
	}
	fmt.Println(string(b.Bytes()))

}
