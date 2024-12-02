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
	dampenedSafeReportCount := 0

	for !parser.EndReached() {
		levels := []int{}
		for {
			level := parser.NextInLine(lib.IntValidator)
			if level == nil {
				break
			}
			levels = append(levels, level.ToInt())
		}

		safe, dampened := assessWithDampening(levels)
		if safe && !dampened {
			safeReportCount++
		}
		if safe {
			dampenedSafeReportCount++
		}

		parser.SkipSentence()
	}

	return safeReportCount, dampenedSafeReportCount, nil
}

func assessWithDampening(levels []int) (safe bool, dampened bool) {
	isSafe, index := assessSafety(levels)
	if isSafe {
		return true, false
	}

	var copiedLevels []int

	copiedLevels = lib.IntCopy(levels)
	levelsDampenedAtStart := copiedLevels[1:]
	isSafe, _ = assessSafety(levelsDampenedAtStart)
	if isSafe {
		return true, true
	}

	copiedLevels = lib.IntCopy(levels)
	levelsDampenedFurther := append(copiedLevels[:index], copiedLevels[index+1:]...)
	isSafe, _ = assessSafety(levelsDampenedFurther)
	if isSafe {
		return true, true
	}

	copiedLevels = lib.IntCopy(levels)
	levelsDampenedFurtherPre := append(copiedLevels[:index-1], copiedLevels[index:]...)
	isSafe, _ = assessSafety(levelsDampenedFurtherPre)
	if isSafe {
		return true, true
	}

	return false, true
}

func assessSafety(levels []int) (bool, int) {
	previousLevel := levels[0]
	trend := "unknown"
	levels = levels[1:]

	for index, level := range levels {
		pos := index + 1

		if trend == "unknown" {
			if level > previousLevel {
				trend = "up"
			}
			if level < previousLevel {
				trend = "down"
			}
		}

		if level < previousLevel && trend == "up" {
			return false, pos
		}
		if level > previousLevel && trend == "down" {
			return false, pos
		}

		dist := lib.IntDistance(level, previousLevel)
		if dist == 0 || dist > 3 {
			return false, pos
		}

		previousLevel = level
	}

	return true, -1
}
