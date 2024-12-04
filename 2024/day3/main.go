package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const InputFileName = "input.txt"

type Operation struct {
	buffer   []rune
	operator string
	arg1     int
	arg2     int
	complete bool
}

func main() {
	data, err := os.ReadFile(InputFileName)
	if err != nil {
		panic(err)
	}

	res1, res2, err := solve(string(data))
	if err != nil {
		panic(err)
	}

	fmt.Println(res1)
	fmt.Println(res2)
}

func solve(data string) (int, int, error) {
	sum := 0

	lines := strings.Split(data, "\n")
	for _, line := range lines {
		operations := extractOps(line)
		for _, op := range operations {
			value, err := op.calc()
			if err == nil {
				sum += value
			}
		}
	}

	return sum, 0, nil
}

func extractOps(line string) (operations []Operation) {
	res := []Operation{}

	current := &Operation{}
	for _, r := range line {
		err := current.ingest(r)
		if err != nil {
			current = &Operation{}
			err = current.ingest(r)
			if err != nil {
				current = &Operation{}
			}
		}

		if current.complete {
			res = append(res, *current)
			current = &Operation{}
		}
	}

	return res
}

func (o *Operation) ingest(r rune) error {
	if o.operator == "" && isLowercaseLetter(r) {
		o.buffer = append(o.buffer, r)
		if len(o.buffer) > 3 {
			o.buffer = o.buffer[1:]
		}
		return nil
	}

	if o.operator == "" && len(o.buffer) == 3 && r == '(' {
		o.operator = string(o.buffer)
		o.buffer = []rune{}
		return nil
	}

	if o.operator != "" && isDigit(r) {
		o.buffer = append(o.buffer, r)
		return nil
	}

	if o.operator != "" && len(o.buffer) > 0 && r == ',' {
		o.arg1, _ = strconv.Atoi(string(o.buffer))
		o.buffer = []rune{}
		return nil
	}

	if o.operator != "" && len(o.buffer) > 0 && r == ')' {
		o.arg2, _ = strconv.Atoi(string(o.buffer))
		o.complete = true
		o.buffer = []rune{}
		return nil
	}

	return fmt.Errorf("unexpected char %c", r)
}

func (o *Operation) calc() (int, error) {
	if o.operator == "mul" {
		return o.arg1 * o.arg2, nil
	}
	return 0, fmt.Errorf("failed to calc %s", o.operator)
}

func isLowercaseLetter(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
