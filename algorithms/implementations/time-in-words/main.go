package main

import (
	"fmt"
	"strconv"
)

func main() {
	var H, M string

	fmt.Scanf("%s\n", &H)
	fmt.Scanf("%s\n", &M)

	m, _ := strconv.Atoi(M)
	h, _ := strconv.Atoi(H)
	var dict = map[string]string{
		"0":  "midnight",
		"1":  "one",
		"2":  "two",
		"3":  "three",
		"4":  "four",
		"5":  "five",
		"6":  "six",
		"7":  "seven",
		"8":  "eight",
		"9":  "nine",
		"10": "ten",
		"11": "eleven",
		"12": "twelve",
		"13": "thirteen",
		"14": "fourteen",
		"15": "fifteen",
		"16": "sixteen",
		"17": "seventeen",
		"18": "eighteen",
		"19": "nineteen",
	}

	var tens = map[int]string{
		2: "twenty",
		3: "thirty",
		4: "forty",
		5: "fifty",
	}
	var answer string
	if m == 0 {
		if H == "0" {
			answer = fmt.Sprintf("midnight")
		} else {
			answer = fmt.Sprintf("%s o' clock", dict[H])
		}
	} else if m < 30 {
		if m%15 == 0 {
			answer = fmt.Sprintf("quarter past %s", dict[H])
		} else {
			answer = fmt.Sprintf("%s past %s", getMinutes(m, dict, tens), dict[H])
		}
	} else if m > 30 {
		if m == 45 {
			answer = fmt.Sprintf("quarter to %s", dict[fmt.Sprintf("%d", h+1)])
		} else {
			answer = fmt.Sprintf("%s to %s", getMinutes(60-m, dict, tens), dict[fmt.Sprintf("%d", h+1)])
		}
	} else {
		// m == 30
		answer = fmt.Sprintf("half past %s", dict[H])
	}
	fmt.Println(answer)
}

func getMinutes(m int, ones map[string]string, seconddig map[int]string) string {
	if m == 1 {
		return "one minute"
	} else if m < 20 {
		return ones[fmt.Sprintf("%d", m)] + " minutes"
	} else {
		var first, second string

		mod := m % 10
		second = ones[fmt.Sprintf("%d", mod)]
		m = m / 10
		mod = m % 10
		first = seconddig[mod]
		return fmt.Sprintf("%s %s minutes", first, second)

	}
}
