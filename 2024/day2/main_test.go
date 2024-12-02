package main

import (
	"fmt"
	"testing"
)

func TestAssessSafety(t *testing.T) {
	tests := []struct {
		levels []int
		safe   bool
		index  int
	}{
		{[]int{7, 6, 4, 2, 1}, true, -1},
		{[]int{1, 2, 7, 8, 9}, false, 2},
		{[]int{9, 7, 6, 2, 1}, false, 3},
		{[]int{1, 3, 2, 4, 5}, false, 2},
		{[]int{8, 6, 4, 4, 1}, false, 3},
		{[]int{1, 3, 6, 7, 9}, true, -1},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			safe, index := assessSafety(test.levels)
			if safe != test.safe {
				t.Errorf("Got safe %t wanted %t", safe, test.safe)
				return
			}
			if index != test.index {
				t.Errorf("Got index %d wanted %d", index, test.index)
				return
			}
		})
	}
}

func TestAssessWithDampening(t *testing.T) {
	tests := []struct {
		levels   []int
		safe     bool
		dampened bool
	}{
		{[]int{7, 6, 4, 2, 1}, true, false},
		{[]int{1, 2, 7, 8, 9}, false, true},
		{[]int{9, 7, 6, 2, 1}, false, true},
		{[]int{1, 3, 2, 4, 5}, true, true},
		{[]int{8, 6, 4, 4, 1}, true, true},
		{[]int{1, 3, 6, 7, 9}, true, false},
		{[]int{2, 1, 2, 3, 4}, true, true}, // dampen at start
		{[]int{4, 1, 5, 7, 9}, true, true}, // dampen before
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			safe, dampened := assessWithDampening(test.levels)
			if safe != test.safe {
				t.Errorf("Got safe %t wanted %t", safe, test.safe)
				return
			}
			if dampened != test.dampened {
				t.Errorf("Got dampened %t wanted %t", dampened, test.dampened)
				return
			}
		})
	}
}
