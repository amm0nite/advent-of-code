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
	for index, leftValue := range leftList {
		rightValue := rightList[index]
		distance := lib.IntDistance(leftValue, rightValue)
		distances = append(distances, distance)
	}

	return lib.IntSum(distances), 0, nil
}
