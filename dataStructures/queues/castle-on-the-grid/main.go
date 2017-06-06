package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// we store the grid in a []string - arr
// each position can be referenc as arr[x][y] -- the character
//
// start at target
// recuresively call a move functin move(current, next, weight, direction)
// current - starts at target
// next - next cell to visit
// weight - starts as 1
// direction - traversal horizontal or vertical
// every time we invoke move with a new direction, increment weight by 1
// if node is visited, and weight of node is greater than new weight coming in, we update the node
// this become DFS i think

// option 2
// move is always called with a m, n+1 and m+1, n and not visiting previously visited node --- BFS

// each node is now not a character but a struct with 3 values, open/closed. x, y, weight

func main() {
	var (
		n        int
		grid     []string
		src_dest string
	)
	fmt.Scanf("%d", &n)
	s := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		s.Scan()
		grid = append(grid, s.Text())
	}
	s.Scan()
	src_dest = s.Text()
	buildnodes(grid)
	co_ord := strings.Split(src_dest, " ")
	//	fmt.Println(nodes)
	//	fmt.Println(co_ord)
	dest := nodes[fmt.Sprintf("%s-%s", co_ord[2], co_ord[3])]
	//	fmt.Println(dest)
	src := nodes[fmt.Sprintf("%s-%s", co_ord[0], co_ord[1])]
	if dest != nil {
		move(dest, src, 0, horizontal)
		move(dest, src, 0, vertical)
	}
	fmt.Println(src.weight)

}

type node struct {
	x, y, weight           int
	visited, allowed, dest bool
}

type direction int

const (
	vertical direction = iota
	horizontal
)

func move(current, dest *node, currentWeight int, d direction) {
	// set the current weight
	//	fmt.Printf("Visiting node %d,%d with weight %d\n", current.x, current.y, currentWeight)
	if !current.allowed {
		return
	}
	if !current.visited || current.weight > currentWeight {
		current.weight = currentWeight
		current.visited = true
	} else {
		// visited this node
		return
	}
	if current.x == dest.x && current.y == dest.y {
		// we reached. lets move on
		return
	}

	current.weight = currentWeight
	nextR := getNode(current.x, current.y+1)
	nextD := getNode(current.x+1, current.y)
	nextL := getNode(current.x, current.y-1)
	nextU := getNode(current.x-1, current.y)

	switch d {
	case horizontal:
		if nextR != nil && nextR.allowed {
			move(nextR, dest, currentWeight, d)
		}
		if nextD != nil && nextD.allowed {
			move(nextD, dest, currentWeight+1, vertical)
		}
		if nextL != nil && nextL.allowed {
			move(nextL, dest, currentWeight, d)
		}
		if nextU != nil && nextU.allowed {
			move(nextU, dest, currentWeight+1, d)
		}

	case vertical:
		if nextR != nil && nextR.allowed {
			move(nextR, dest, currentWeight+1, d)
		}
		if nextD != nil && nextD.allowed {
			move(nextD, dest, currentWeight, vertical)
		}
		if nextL != nil && nextL.allowed {
			move(nextL, dest, currentWeight+1, d)
		}
		if nextU != nil && nextU.allowed {
			move(nextU, dest, currentWeight, d)
		}

	}
}

func buildnodes(vals []string) {
	for i, v := range vals {
		for j := 0; j < len(v); j++ {
			var allowed bool
			if v[j] == 'X' {
				allowed = false
			} else {
				allowed = true
			}
			n := &node{x: i, y: j, allowed: allowed, weight: 0}
			nodes[fmt.Sprintf("%d-%d", i, j)] = n
		}
	}
}

var nodes = make(map[string]*node)

func getNode(x, y int) *node {
	return nodes[fmt.Sprintf("%d-%d", x, y)]

}
