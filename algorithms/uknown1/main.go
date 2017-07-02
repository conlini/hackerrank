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
	temp := strings.Split(s.Text(), " ")
	m, _ := strconv.Atoi(temp[0])
	w, _ := strconv.Atoi(temp[1])
	p, _ := strconv.Atoi(temp[2])
	n, _ := strconv.Atoi(temp[3])
	fmt.Println(getPasses(m, w, p, n))
}

func getPasses(m, w, p, n int) int {
	var passes int
	passes++
	var budget = m * w
	for budget != n {
		count := budget / p
		//		fmt.Println(m, w, budget, count, passes)
		m, w, budget = findBest(m, w, count, n)
		//		fmt.Println(m, w, budget)
		passes++

	}
	return passes
}

func findBest(ma, wo, count, n int) (int, int, int) {
	if ma < wo {
		ma += count
		if ma*wo > n {
			ma--
		}

	} else if ma > wo {
		wo += count
		if ma*wo > n {
			wo--
		}

	}
	if ma == wo {
		ma--
	}
	return ma, wo, ma * wo
}
