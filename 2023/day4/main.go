package main

import (
	"advent-of-code/lib"
	"fmt"
	"math"
	"os"
)

const InputFileName = "input.txt"

func main() {
	data, err := os.ReadFile(InputFileName)
	if err != nil {
		panic(err)
	}

	res, err := solve(string(data))
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

type Pile struct {
	cards []*Card
}

type Card struct {
	number  int
	winning []int
	roster  []int
}

func IsInt(t *lib.Token) bool {
	return t.IsInt()
}

func solve(raw string) (int, error) {
	parser := lib.CreateParser(raw, []rune{':', ' ', '|'})
	pile := &Pile{}

	for {
		card := &Card{}

		n := parser.NextInLine(IsInt)
		if n == nil {
			break
		}
		card.number = n.ToInt()

		for {
			w := parser.NextInLine(func(t *lib.Token) bool { return t.IsInt() || t.ToRune() == '|' })
			if !w.IsInt() {
				break
			}
			card.winning = append(card.winning, w.ToInt())
		}
		for {
			r := parser.NextInLine(IsInt)
			if r == nil {
				break
			}
			card.roster = append(card.roster, r.ToInt())
		}

		pile.cards = append(pile.cards, card)
		parser.SkipSentence()
	}

	return pile.solve1(), nil
}

func (c *Card) score() int {
	sum := 0
	for _, w := range c.winning {
		if lib.IntSliceContains(c.roster, w) {
			sum++
		}
	}

	score := int(math.Pow(2, float64(sum-1)))
	fmt.Println(c.winning, c.roster, sum, score)
	return score
}

func (p *Pile) solve1() int {
	sum := 0
	for _, c := range p.cards {
		sum += c.score()
	}
	return sum
}

func (c *Card) String() string {
	return fmt.Sprintf("(#%d %d %d)", c.number, len(c.winning), len(c.roster))
}
