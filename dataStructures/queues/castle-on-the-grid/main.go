package main

import (
	"bufio"
	"bytes"
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
	//	//fmt.Println(nodes)
	//	//fmt.Println(co_ord)
	dest := nodes[fmt.Sprintf("%s-%s", co_ord[2], co_ord[3])]
	//	//fmt.Println(dest)
	src := nodes[fmt.Sprintf("%s-%s", co_ord[0], co_ord[1])]
	fmt.Println(movebfs(src, dest, 0, n, vertical))

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

func print(q []*node) {
	buf := new(bytes.Buffer)
	for _, n := range q {
		buf.Write([]byte(fmt.Sprintf("%+v", n)))
		buf.Write([]byte(","))
	}
	//fmt.Println(string(buf.Bytes()))
}
func movebfs(src, dest *node, currentWeight, n int, d direction) int {
	// lets have a queue.. since everyone wants it
	q := make([]*node, 0, len(nodes))
	src.weight = 0
	q = append(q, src)
	for len(q) > 0 {
		//fmt.Println("***************")
		//print(q)
		// get top
		top := q[0]
		if top.x == dest.x && top.y == dest.y {
			return top.weight
		}
		// re-adust q
		q = q[1:]
		//		top.weight = currentWeigh
		// find all east nodes. They will have a weight of 1
		y := top.y
		x := top.x
		for y < n-1 {
			y++
			node := getNode(x, y)
			//fmt.Println("east node - ", node)
			if node != nil && !node.visited {
				if node.allowed {
					//fmt.Println("adding east node", node)
					q = appendToq(q, node, top.weight)
				} else {
					break
				}
			}

		}
		// find all west nodes. They will have a weight of 1
		y = top.y
		x = top.x
		for y > 0 {
			y--
			node := getNode(x, y)
			//fmt.Println("west node - ", node)
			if node != nil && !node.visited {
				if node.allowed {
					//fmt.Println("adding westt node", node)
					q = appendToq(q, node, top.weight)
				} else {
					break
				}
			}

		}
		// find all north nodes. They will have a weight of 1
		y = top.y
		x = top.x
		for x > 0 {
			x--
			node := getNode(x, y)
			//fmt.Println("north node - ", node)
			if node != nil && !node.visited {
				if node.allowed {
					//fmt.Println("adding north node", node)
					q = appendToq(q, node, top.weight)
				} else {
					break
				}
			}

		}
		// find all south nodes. They will have a weight of 1
		y = top.y
		x = top.x
		for x < n-1 {
			x++
			node := getNode(x, y)
			//fmt.Println("south node - ", node)
			if node != nil && !node.visited {
				if node.allowed {
					//fmt.Println("adding southt node", node)
					q = appendToq(q, node, top.weight)
				} else {
					break
				}
			}

		}

	}
	return -1
}

func appendToq(q []*node, node *node, weight int) []*node {
	node.weight = weight + 1
	node.visited = true
	q = append(q, node)
	return q
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
