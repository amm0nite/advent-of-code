package main

import (
	"advent-of-code/lib"
	"fmt"
	"os"
)

const InputFileName = "input.txt"

func isPart(t *lib.Token) bool {
	return !t.IsInt() && t.Buffer != "."
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

type PartSymbol struct {
	line   int
	col    int
	symbol rune
}

type PartNumber struct {
	line  int
	start int
	end   int
	value int
}

func solve(input string) (int, int, error) {
	list := ".*#$&=+-%@/"
	symbols := []rune{}
	for _, c := range list {
		symbols = append(symbols, c)
	}

	parser := lib.CreateParser(input, symbols)

	parser.ResetCursor()
	partSymbols := []PartSymbol{}

	for {
		sym := parser.Next(isPart)
		if sym == nil {
			break
		}
		partSymbol := &PartSymbol{line: sym.Line, col: sym.Col, symbol: sym.ToRune()}
		partSymbols = append(partSymbols, *partSymbol)
	}

	parser.ResetCursor()
	partNumbers := []PartNumber{}

	for {
		number := parser.Next(func(t *lib.Token) bool { return t.IsInt() })
		if number == nil {
			break
		}
		partNumber := &PartNumber{line: number.Line, start: number.Col, end: number.Col + number.Len() - 1, value: number.ToInt()}
		partNumbers = append(partNumbers, *partNumber)
	}

	fmt.Println(partSymbols)
	fmt.Println(partNumbers)

	schematic := &Schematic{partSymbols: partSymbols, partNumbers: partNumbers}
	return schematic.solve1(), schematic.solve2(), nil
}

type Schematic struct {
	partSymbols []PartSymbol
	partNumbers []PartNumber
}

func isAdjacent(ps PartSymbol, pn PartNumber) bool {
	x0 := ps.col
	y0 := ps.line

	x1 := pn.start
	y1 := pn.line

	x2 := pn.end
	y2 := pn.line

	if lib.IsNeighbour(x0, y0, x1, y1) {
		return true
	}
	if lib.IsNeighbour(x0, y0, x2, y2) {
		return true
	}
	return false
}

func (s *Schematic) solve1() int {
	selection := []int{}
	for _, partSymbol := range s.partSymbols {
		for _, partNumber := range s.partNumbers {
			if isAdjacent(partSymbol, partNumber) {
				selection = append(selection, partNumber.value)
			}
		}
	}

	return lib.IntSum(selection)
}

func (s *Schematic) solve2() int {
	ratios := []int{}

	for _, partSymbol := range s.partSymbols {
		if partSymbol.symbol != '*' {
			continue
		}

		numbers := []int{}
		for _, partNumber := range s.partNumbers {
			if isAdjacent(partSymbol, partNumber) {
				numbers = append(numbers, partNumber.value)
			}
		}

		if len(numbers) == 2 {
			ratio := lib.IntProduct(numbers)
			ratios = append(ratios, ratio)
		}
	}

	return lib.IntSum(ratios)
}
