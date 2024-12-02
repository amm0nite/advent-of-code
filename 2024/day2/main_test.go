package main

import (
	"fmt"
	"testing"
)

func TestAssessSafety(t *testing.T) {
	tests := []struct {
		levels []int
		safe   bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, false},
		{[]int{8, 6, 4, 4, 1}, false},
		{[]int{1, 3, 6, 7, 9}, true},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			res := assessSafety(test.levels)
			if res != test.safe {
				t.Errorf("Got %t wanted %t", res, test.safe)
				return
			}
		})
	}
}
