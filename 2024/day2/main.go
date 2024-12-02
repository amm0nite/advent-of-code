package main

import (
	"advent-of-code/lib"
	"fmt"
	"os"
)

const InputFileName = "input.txt"

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
	parser := lib.CreateParser(data, []rune{' '})

	safeReportCount := 0

	for !parser.EndReached() {
		levels := []int{}
		for {
			level := parser.NextInLine(lib.IntValidator)
			if level == nil {
				break
			}
			levels = append(levels, level.ToInt())
		}

		if assessSafety(levels) {
			safeReportCount++
		}

		parser.SkipSentence()
	}

	return safeReportCount, 0, nil
}

func assessSafety(levels []int) bool {
	previousLevel := levels[0]
	trend := "unknown"
	levels = levels[1:]

	for _, level := range levels {
		if trend == "unknown" {
			if level > previousLevel {
				trend = "up"
			}
			if level < previousLevel {
				trend = "down"
			}
		}

		if level < previousLevel && trend == "up" {
			return false
		}
		if level > previousLevel && trend == "down" {
			return false
		}

		dist := lib.IntDistance(level, previousLevel)
		if dist == 0 || dist > 3 {
			return false
		}

		previousLevel = level
	}

	return true
}
