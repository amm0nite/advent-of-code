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
	filteredSum := 0
	enabled := true

	lines := strings.Split(data, "\n")
	for _, line := range lines {
		operations := extractOps(line)
		for _, op := range operations {
			if op.operator == "don't" {
				enabled = false
			}
			if op.operator == "do" {
				enabled = true
			}

			value, err := op.calc()
			if err == nil {
				sum += value
				if enabled {
					filteredSum += value
				}
			}
		}
	}

	return sum, filteredSum, nil
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
	if o.operator == "" && (isLowercaseLetter(r) || r == '\'') {
		o.buffer = append(o.buffer, r)
		return nil
	}

	if o.operator == "" && len(o.buffer) > 1 && r == '(' {
		o.operator = cleanOperator(string(o.buffer))
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

	if o.operator != "" && len(o.buffer) == 0 && r == ')' {
		o.complete = true
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

func cleanOperator(operator string) string {
	if strings.HasSuffix(operator, "mul") {
		return "mul"
	}
	if strings.HasSuffix(operator, "do") {
		return "do"
	}
	if strings.HasSuffix(operator, "don't") {
		return "don't"
	}
	return operator
}

func isLowercaseLetter(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
