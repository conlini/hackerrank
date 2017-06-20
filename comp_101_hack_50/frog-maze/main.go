package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cell struct {
	val         string
	i2, j2      int
	i1, j1      int
	probability float64
	visited     bool
}

func (c cell) String() string { return fmt.Sprintf("{Val - %s, i = %d, j = %d}", c.val, c.i1, c.j1) }

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	var n, m, k int
	f := strings.Split(s.Text(), " ")
	n, _ = strconv.Atoi(f[0])
	m, _ = strconv.Atoi(f[1])
	k, _ = strconv.Atoi(f[2])

	var alex cell
	maze := make([][]cell, n)
	for i := 0; i < n; i++ {
		s.Scan()
		row := make([]cell, m)
		split := strings.Split(s.Text(), "")
		for j, v := range split {
			row[j] = cell{val: v, i1: i, j1: j, probability: 1, i2: -1, j2: -1}
			if v == "A" {
				alex = row[j]
			}
		}
		maze[i] = row
	}
	for i := 0; i < k; i++ {
		s.Scan()
		f := strings.Split(s.Text(), " ")
		i1, _ := strconv.Atoi(f[0])
		j1, _ := strconv.Atoi(f[1])
		i2, _ := strconv.Atoi(f[2])
		j2, _ := strconv.Atoi(f[3])
		maze[i1][j1].i2 = i2
		maze[i1][j1].j2 = j2
	}
	//	fmt.Println("starting walk")
	q := make([]cell, 0)
	q = append(q, alex)
	for len(q) > 0 {
		//	fmt.Println("q-->", q)

		current := q[0]
		if current.val == "%" {
			fmt.Println(current.probability)
			return
		}
		next := make([]cell, 0)
		if current.i1-1 >= 0 {
			c := maze[current.i1-1][current.j1]
			if c.val != "#" {
				next = append(next, c)
			}
		}
		if current.i1+1 < n {
			c := maze[current.i1+1][current.j1]
			if c.val != "#" {
				next = append(next, c)
			}
		}
		if current.j1-1 >= 0 {
			c := maze[current.i1][current.j1-1]
			if c.val != "#" {
				next = append(next, c)
			}
		}
		if current.j1+1 < m {
			c := maze[current.i1][current.j1+1]
			if c.val != "#" {
				next = append(next, c)
			}
		}
		//	fmt.Println("next -->", next)
		if len(next) > 0 {
			prob := 1.0 / float64(len(next))
			//		fmt.Println("probability -->", prob)
			for _, v := range next {
				if !v.visited {
					v.visited = true
					if v.val != "*" {
						if v.j2 != -1 && v.i2 != -1 {
							jump := maze[v.i2][v.j2]
							jump.probability = jump.probability * prob
							jump.visited = true

							q = append(q, jump)
						} else {
							v.probability = v.probability * prob
							q = append(q, v)
						}
					}
				}

			}
		}
		q = q[1:]
	}

}
