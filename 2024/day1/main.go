package main

import (
	"advent-of-code/lib"
	"fmt"
	"os"
	"sort"
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

	leftList := []int{}
	rightList := []int{}

	parser := lib.CreateParser(data, []rune{' '})

	for {
		leftN := parser.NextInLine(lib.IntValidator)
		if leftN == nil {
			break
		}
		leftList = append(leftList, leftN.ToInt())
		rightN := parser.NextInLine(lib.IntValidator)
		if rightN == nil {
			break
		}
		rightList = append(rightList, rightN.ToInt())

		parser.SkipSentence()
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	distances := []int{}
	similarities := []int{}

	for index, leftValue := range leftList {
		// part1
		rightValue := rightList[index]
		distance := lib.IntDistance(leftValue, rightValue)
		distances = append(distances, distance)

		// part2
		count := lib.IntCountOccurences(rightList, leftValue)
		similarities = append(similarities, leftValue*count)
	}

	return lib.IntSum(distances), lib.IntSum(similarities), nil
}
