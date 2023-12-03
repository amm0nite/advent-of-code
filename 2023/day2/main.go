package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const InputFileName = "input.txt"

const MaxRed = 12
const MaxGreen = 13
const MaxBlue = 14

var Separators = []rune{' ', ',', ';', ':'}
var Colors = []string{"red", "green", "blue"}

type Game struct {
	selection []int
}

type Token struct {
	buffer string
}

type Sentence struct {
	cursor int
	tokens []Token
}

func (t *Token) isInt() bool {
	_, err := strconv.Atoi(t.buffer)
	return err == nil
}

func (t *Token) toInt() int {
	value, err := strconv.Atoi(t.buffer)
	if err != nil {
		return 0
	}
	return value
}

func (t *Token) isEmpty() bool {
	return t.buffer == ""
}

func (t *Token) append(c rune) {
	t.buffer += string(c)
}

func (t *Token) isColor() bool {
	for _, c := range Colors {
		if t.buffer == c {
			return true
		}
	}
	return false
}

func (s *Sentence) nextColor() (*Token, error) {
	for s.cursor < len(s.tokens) {
		i := s.cursor
		s.cursor++

		if s.tokens[i].isColor() {
			return &s.tokens[i], nil
		}

	}
	return nil, fmt.Errorf("color not found")
}

func (s *Sentence) nextInt() (*Token, error) {
	for s.cursor < len(s.tokens) {
		i := s.cursor
		s.cursor++

		if s.tokens[i].isInt() {
			return &s.tokens[i], nil
		}
	}
	return nil, fmt.Errorf("int not found")
}

func getLines() ([]string, error) {
	data, err := os.ReadFile(InputFileName)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}

func main() {
	lines, err := getLines()
	if err != nil {
		panic(err)
	}

	game := &Game{}
	err = game.solve(lines)
	if err != nil {
		panic(err)
	}

	fmt.Println(game.res())
}

func isSeparator(c rune) bool {
	for _, s := range Separators {
		if c == s {
			return true
		}
	}
	return false
}

func parse(line string) *Sentence {
	tokens := []Token{}
	current := &Token{}

	for _, c := range line {
		if isSeparator(c) {
			if !current.isEmpty() {
				tokens = append(tokens, *current)
				current = &Token{}
			}
			current.append(c)
			tokens = append(tokens, *current)
			current = &Token{}
			continue
		}

		current.append(c)
	}

	if !current.isEmpty() {
		tokens = append(tokens, *current)
	}

	return &Sentence{tokens: tokens}
}

func (g *Game) solve(lines []string) error {
	for _, line := range lines {
		if line == "" {
			continue
		}

		sentence := parse(line)
		err := g.process(sentence)
		if err != nil {
			return err
		}
	}

	return nil
}

type BagHint struct {
	red   int
	green int
	blue  int
}

func (bh *BagHint) String() string {
	return fmt.Sprintf("r:%d,g:%d,b:%d", bh.red, bh.green, bh.blue)
}

func (bh *BagHint) valid() bool {
	return bh.red <= MaxRed && bh.green <= MaxGreen && bh.blue <= MaxBlue
}

func intMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func (g *Game) process(sentence *Sentence) error {
	id, err := sentence.nextInt()
	if err != nil {
		return err
	}

	hint := &BagHint{}
	for {
		count, err := sentence.nextInt()
		if err != nil {
			break
		}
		color, err := sentence.nextColor()
		if err != nil {
			break
		}

		if color.buffer == "red" {
			hint.red = intMax(hint.red, count.toInt())
		}
		if color.buffer == "green" {
			hint.green = intMax(hint.green, count.toInt())
		}
		if color.buffer == "blue" {
			hint.blue = intMax(hint.blue, count.toInt())
		}
	}

	if hint.valid() {
		g.selection = append(g.selection, id.toInt())
	}

	return nil
}

func (g *Game) res() int {
	sum := 0
	for _, s := range g.selection {
		sum += s
	}
	return sum
}
