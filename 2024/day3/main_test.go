package main

import (
	"fmt"
	"testing"
)

func TestExtractOps(t *testing.T) {
	tests := []struct {
		line       string
		operations []Operation
	}{
		{
			"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			[]Operation{
				{operator: "mul", arg1: 2, arg2: 4},
				{operator: "mul", arg1: 5, arg2: 5},
				{operator: "mul", arg1: 11, arg2: 8},
				{operator: "mul", arg1: 8, arg2: 5},
			},
		},
		{
			"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			[]Operation{
				{operator: "mul", arg1: 2, arg2: 4},
				{operator: "don't", arg1: 0, arg2: 0},
				{operator: "mul", arg1: 5, arg2: 5},
				{operator: "mul", arg1: 11, arg2: 8},
				{operator: "do", arg1: 0, arg2: 0},
				{operator: "mul", arg1: 8, arg2: 5},
			},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ops := extractOps(test.line)
			fmt.Println(ops)
			if len(ops) != len(test.operations) {
				t.Errorf("Got %d ops, want %d", len(ops), len(test.operations))
				return
			}

			for index, op := range ops {
				expected := test.operations[index]
				if op.operator != expected.operator {
					t.Errorf("Got %s operator, want %s", op.operator, expected.operator)
					return
				}
				if op.arg1 != expected.arg1 {
					t.Errorf("Got %d arg1, want %d", op.arg1, expected.arg1)
					return
				}
				if op.arg2 != expected.arg2 {
					t.Errorf("Got %d arg2, want %d", op.arg2, expected.arg2)
					return
				}
			}
		})
	}
}
