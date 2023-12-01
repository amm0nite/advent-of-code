package main

import "testing"

func TestSolve(t *testing.T) {
	res, err := solve([]string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"})
	if err != nil {
		t.Error(err)
	}
	expected := 142
	if res != 142 {
		t.Errorf("res = %d; want %d", res, expected)
	}
}
