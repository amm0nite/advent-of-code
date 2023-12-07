package main

import (
	"advent-of-code/lib"
	"fmt"
)

const InputFileName = "input_test.txt"

func isPart(t *lib.Token) bool {
	return !t.IsInt() && t.Buffer != "."
}

func main() {
	res, err := solve(InputFileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

type PartSymbol struct {
	line int
	col  int
}

type PartNumber struct {
	line  int
	start int
	end   int
	value int
}

func solve(filename string) (int, error) {
	list := ".*#$&=+-%@/"
	symbols := []rune{}
	for _, c := range list {
		symbols = append(symbols, c)
	}

	parser, err := lib.CreateParser(filename, symbols)
	if err != nil {
		panic(err)
	}

	parser.ResetCursor()
	partSymbols := []PartSymbol{}

	for {
		sym := parser.Next(isPart)
		if sym == nil {
			break
		}
		partSymbol := &PartSymbol{line: sym.Line, col: sym.Col}
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

	selection := &lib.IntSet{}
	for _, partSymbol := range partSymbols {
		for _, partNumber := range partNumbers {
			x0 := partSymbol.col
			y0 := partSymbol.line

			x1 := partNumber.start
			y1 := partNumber.line

			x2 := partNumber.end
			y2 := partNumber.line

			if lib.IsNeighbour(x0, y0, x1, y1) {
				selection.Add(partNumber.value)
				continue
			}
			if lib.IsNeighbour(x0, y0, x2, y2) {
				selection.Add(partNumber.value)
				continue
			}
		}
	}

	fmt.Println(selection)
	return selection.Sum(), nil
}
