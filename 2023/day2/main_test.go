package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		lines []string
		res1  int
		res2  int
	}{
		{
			[]string{
				"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
				"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
				"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
				"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
				"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
			8,
			2286,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			game := &Game{}
			game.solve(test.lines)
			res1 := game.res1()
			res2 := game.res2()
			if res1 != test.res1 {
				t.Errorf("res1=%d want %d", res1, test.res1)
			}
			if res2 != test.res2 {
				t.Errorf("res2=%d want %d", res2, test.res2)
			}
		})
	}
}
