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
		calibration, err := CreateCalibrationFromLine(line)
		if err != nil {
			return 0, err
		}
		doc.values = append(doc.values, calibration)
	}

	return doc.sum(), nil
}

func CreateCalibrationFromLine(line string) (*Calibration, error) {
	r, err := regexp.Compile("([0-9])")
	if err != nil {
		return nil, err
	}

	findings := r.FindAllStringSubmatch(line, -1)
	digits := []int{}

	for _, f := range findings {
		intval, err := strconv.Atoi(f[len(f)-1])
		if err != nil {
			panic(err)
		}

		digits = append(digits, intval)
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
