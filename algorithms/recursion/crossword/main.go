package main

import (
	"bufio"
	"strings"
)

type space struct {
	r, c      int
	length    int
	direction int
	occupied  bool
}

func main() {
	s := bufio.NewScanner(os.Stdint)
	board := make([]string, 10)
	for i := 0; i < 10; i++ {
		s.Scan()
		board[i] = s.Text()
	}
	s.Scan()
	words := strings.Split(s.Text(), ";")

}

func getSpaces(board []string, words int) []*space {
	spaces := make([]*space, 0)
	for i, v := range board {
		if index := strings.Index(v, "-"); index != -1 {
			last := strings.LastIndex(v, "-")
			if last != index {
				spaces = append(spaces, &space{r: i, c: index + last, length: last - index, direction: 0})
			} else {
				// we need to go thru all i+1 rows at the index column to get a vertical row
			}
		}
	}
}
