package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	expected1 := 4361
	expected2 := 467835

	res1, res2, err := solve(strings.Join(input, "\n"))
	if err != nil {
		t.Error(err)
	}
	if res1 != expected1 {
		t.Errorf("res1=%d want %d", res1, expected1)
	}
	if res2 != expected2 {
		t.Errorf("res2=%d want %d", res2, expected2)
	}
}
