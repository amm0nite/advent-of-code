package main

import (
	"advent-of-code/lib"
	"fmt"
)

const InputFileName = "input.txt"

func isPart(t *lib.Token) bool {
	return !t.IsInt() && t.Buffer != "."
}

func main() {
	res, err := solve()
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
	col   int
	value int
}

func solve() (int, error) {
	list := ".*#$&=+-%@/"
	symbols := []rune{}
	for _, c := range list {
		symbols = append(symbols, c)
	}

	parser, err := lib.CreateParser(InputFileName, symbols)
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
		partNumber := &PartNumber{line: number.Line, col: number.Col, value: number.ToInt()}
		partNumbers = append(partNumbers, *partNumber)
	}

	fmt.Println(partSymbols)
	fmt.Println(partNumbers)

	return 0, nil
}
