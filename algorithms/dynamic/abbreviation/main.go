package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	q, _ := strconv.Atoi(s.Text())
	for i := 0; i < q; i++ {
		s.Scan()
		a := s.Text()
		s.Scan()
		b := s.Text()
		la := len(a)
		lb := len(b)
		dp := make([][]bool, lb+1)
		dp = presetDp(dp, a, b)
		//		fmt.Println(dp)
		for j := 1; j < lb+1; j++ {

			var bchar, achar string
			bchar = string(b[j-1])

			for k := 1; k < la+1; k++ {
				achar = string(a[k-1])

				if bchar == achar || bchar == string([]byte(achar)[0]-32) {
					dp[j][k] = (true && dp[j-1][k-1])
				} else {
					if isCapital(achar) {
						dp[j][k] = false
					} else if isLower(achar) {
						dp[j][k] = dp[j][k-1]
					}
				}
			}
		}
		if dp[lb][la] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func presetDp(dp [][]bool, a string, b string) [][]bool {
	for i := 0; i < len(b)+1; i++ {
		arr := make([]bool, len(a)+1)
		dp[i] = arr
		if i == 0 {
			dp[i][0] = true
		} else {
			dp[i][0] = false
		}
	}

	for i := 1; i < len(a)+1; i++ {
		//		if isLower(string(a[i-1])) {
		//			dp[0][i] = true && dp[0][i-1]
		//		} else {
		//			dp[0][i] = false
		//		}
		dp[0][i] = true
	}
	return dp
}

func isCapital(achar string) bool {
	charAt := []byte(achar)[0]
	return charAt >= 65 && charAt <= 90

}

func isLower(achar string) bool {
	charAt := []byte(achar)[0]
	return charAt >= 97 && charAt <= 122
}
