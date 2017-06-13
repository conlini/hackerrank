package main

import "testing"

func TestMinDiff(t *testing.T) {
	testData := []struct {
		in   []int
		diff int
	}{
		{[]int{1, 99, -9, 10, 3, 44}, 2},
		{[]int{1, 99, -9, 10, 100, 44}, 8},
	}
	for i, td := range testData {
		answer := minAbsDiff(td.in)
		if answer != td.diff {
			t.Errorf("Test %d. Input %v. Expected %d but got %d", i, td.in, td.diff, answer)
			t.FailNow()
		}
	}
}
