package main

import (
	"fmt"
	"testing"
)

func TestCreateCalibrationFromLine(t *testing.T) {
	tests := []struct {
		line     string
		expected int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
		{"eighthree", 83},
		{"sevenine", 79},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("test:%s", test.line), func(t *testing.T) {
			res, err := createCalibrationFromLine(test.line)
			if err != nil {
				t.Error(err)
			}
			if res.intval() != test.expected {
				t.Errorf("%d %d (%d) want %d", res.first, res.second, res.intval(), test.expected)
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
