package main

import (
	"bufio"
	"fmt"
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

func getSpaces(board []string, words int) map[int][]*space {
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

func check(words, board []string, spaces map[int][]*space) {
	if len(words) == 0 {
		fmt.Println(board)
		return
	}
	// get next word
	next := words[0]
	remaining := words[1:]
	// get next available spaces
	possibleSlots := spaces[len(next)]

	for _, s := range possibleSlots {
		if !s.occupied {
			if b, OK := canFit(word, s, board); OK {
				s.occupied = true
				check(remaining, b, spaces)
			}
		}
	}

}

func canFit(word string, s *space, board []string) ([]string, OK) {
	return nil, nil
}
