package main

import (
	"fmt"
	"testing"
)

func TestExtractDigits(t *testing.T) {
	tests := []struct {
		line     string
		expected []int
	}{
		{"treb7uchet", []int{7}},
		{"two1nine", []int{2, 1, 9}},
		{"eightwothree", []int{8, 2, 3}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("test:%s", test.line), func(t *testing.T) {
			res, err := extractDigitsFromLine(test.line)
			if err != nil {
				t.Error(err)
			}
			if len(res) != len(test.expected) {
				t.Errorf("len(res) = %d; want %d", len(res), len(test.expected))
			}
			for i, v := range res {
				if v != test.expected[i] {
					t.Errorf("res[%d] = %d; want %d", i, v, test.expected[i])
				}
			}
		})
	}
}

func TestSolve(t *testing.T) {
	tests := []struct {
		lines    []string
		expected int
	}{
		{[]string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}, 142},
		{[]string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}, 281},
	}

	for index, test := range tests {
		t.Run(fmt.Sprintf("test%d", index), func(t *testing.T) {
			res, err := solve(test.lines)
			if err != nil {
				t.Error(err)
			}
			if res != test.expected {
				t.Errorf("res = %d; want %d", res, test.expected)
			}
		})
	}
}
