package main

import "testing"

func TestIsFunny(t *testing.T) {
	data := []struct{ data, output string }{
		{data: "acxz", output: "Funny"},
		{data: "bcxz", output: "Not Funny"},
		{data: "aaaa", output: "Funny"},
		{data: "acdefggfedca", output: "Funny"},
		{data: "", output: "Funny"},
	}
	for i, test := range data {
		answer := isFunny(test.data)
		if answer != test.output {
			t.Errorf("Test %d. Input %s. Expected %s but got %s", i, test.data, test.output, answer)
			t.FailNow()
		}
	}
}
