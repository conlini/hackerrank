package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cell struct {
	val     string
	r, c    int
	count   int
	visited bool
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	t, _ := strconv.Atoi(s.Text())

	for i := 0; i < t; i++ {
		s.Scan()
		temp := strings.Split(s.Text(), " ")

		n, _ := strconv.Atoi(temp[0])
		m, _ := strconv.Atoi(temp[1])
		maze := make([][]*cell, n)
		var start *cell
		for j := 0; j < n; j++ {
			s.Scan()
			row := make([]*cell, m)
			for k, v := range strings.Split(s.Text(), "") {
				//	fmt.Println(k, v, m, n)
				row[k] = &cell{val: v, r: j, c: k, count: 0}
				if v == "M" {
					start = row[k]
				}
			}
			maze[j] = row
		}
		s.Scan()
		k, _ := strconv.Atoi(s.Text())

		stack := make([]*cell, 0, 4)
		stack = append(stack, start)
		var current *cell
		var count int
		for len(stack) > 0 {
			// get 0'th element
			//			fmt.Println(len(stack))
			current = stack[0]
			current.visited = true
			stack = stack[1:]
			count = current.count
			if current.val == "*" {
				if count == k {
					fmt.Println("Impressed")
				} else {
					fmt.Println("Oops!")
				}
				return
			}
			next := getNext(current, maze, m, n)
			if len(next) > 1 {
				count = count + 1
			}
			//			fmt.Println("next-", next)
			for _, c := range next {

				if !c.visited {
					c.visited = true
					c.count = count
					stack = append([]*cell{c}, stack[0:]...)

				}
				//				fmt.Println("stack-", stack, " c-", c)
			}

		}
	}
}

func (c *cell) String() string {
	return fmt.Sprintf("{val:%s, r:%d, c:%d, count:%d}", c.val, c.r, c.c, c.count)
}

func getNext(current *cell, maze [][]*cell, m, n int) []*cell {
	type pos struct {
		r, c int
	}
	//	fmt.Println(maze)
	answer := make([]*cell, 0)
	index := []pos{pos{current.r - 1, current.c}, pos{current.r + 1, current.c}, pos{current.r, current.c - 1}, pos{current.r, current.c + 1}}
	//	fmt.Println(index)

	if index[0].r >= 0 && maze[index[0].r][index[0].c].val != "X" && !maze[index[0].r][index[0].c].visited {
		answer = append(answer, maze[index[0].r][index[0].c])
	}
	if index[1].r < n && maze[index[1].r][index[1].c].val != "X" && !maze[index[1].r][index[1].c].visited {
		answer = append(answer, maze[index[1].r][index[1].c])
	}
	if index[2].c >= 0 && maze[index[2].r][index[2].c].val != "X" && !maze[index[2].r][index[2].c].visited {
		answer = append(answer, maze[index[2].r][index[2].c])
	}
	if index[3].c < m && maze[index[3].r][index[3].c].val != "X" && !maze[index[3].r][index[3].c].visited {
		answer = append(answer, maze[index[3].r][index[3].c])
	}

	return answer
}
