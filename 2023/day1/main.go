package main

import (
	"fmt"
	"os"
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

func wordToDigit(word string) int {
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for index, number := range numbers {
		if strings.HasSuffix(word, number) {
			return index + 1
		}
	}
	return 0
}

func extractDigitsFromLine(line string) ([]int, error) {
	digits := []int{}
	buffer := ""

	for _, c := range line {
		number, err := strconv.Atoi(string(c))
		if err != nil {
			buffer += string(c)
			number = wordToDigit(buffer)
			if number < 1 {
				continue
			}
			buffer = buffer[len(buffer)-1:]
		}
		digits = append(digits, number)
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
