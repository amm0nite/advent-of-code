package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const InputFileName = "input.txt"

type Document struct {
	values []*Calibration
}

type Calibration struct {
	first  int
	second int
}

func main() {
	data, err := os.ReadFile(InputFileName)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	res, err := solve(lines)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func solve(lines []string) (int, error) {
	doc := &Document{}

	for _, line := range lines {
		if line == "" {
			continue
		}
		calibration, err := createCalibrationFromLine(line)
		if err != nil {
			return 0, err
		}
		doc.values = append(doc.values, calibration)
	}

	return doc.sum(), nil
}

func unspellDigits(line string) string {
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	modifiedLine := line
	for i, n := range numbers {
		value := i + 1
		r := regexp.MustCompile(n)
		modifiedLine = r.ReplaceAllString(modifiedLine, strconv.Itoa(value))
	}

	return modifiedLine
}

func extractDigitsFromLine(line string) ([]int, error) {
	line = unspellDigits(line)

	r, err := regexp.Compile("([0-9])")
	if err != nil {
		return nil, err
	}

	findings := r.FindAllStringSubmatch(line, -1)
	digits := []int{}

	for _, f := range findings {
		found := f[len(f)-1]

		if len(found) == 1 {
			intval, err := strconv.Atoi(f[len(f)-1])
			if err != nil {
				return nil, err
			}
			digits = append(digits, intval)
		}
	}

	return digits, nil
}

func createCalibrationFromLine(line string) (*Calibration, error) {
	digits, err := extractDigitsFromLine(line)
	if err != nil {
		return nil, err
	}

	if len(digits) == 0 {
		digits = append(digits, 0)
	}
	if len(digits) == 1 {
		digits = append(digits, digits[0])
	}

	calibration := &Calibration{first: digits[0], second: digits[len(digits)-1]}
	return calibration, nil
}

func (c *Calibration) intval() int {
	return c.first*10 + c.second
}

func (d *Document) sum() int {
	sum := 0
	for _, v := range d.values {
		sum += v.intval()
	}
	return sum
}
