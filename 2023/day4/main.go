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

	res1, res2, err := solve(string(data))
	if err != nil {
		panic(err)
	}

	fmt.Println(res1)
	fmt.Println(res2)
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

func solve(raw string) (int, int, error) {
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

	return pile.solve1(), pile.solve2(), nil
}

func (c *Card) matches() int {
	sum := 0
	for _, w := range c.winning {
		if lib.IntSliceContains(c.roster, w) {
			sum++
		}
	}
	return sum
}

func (c *Card) score() int {
	matches := c.matches()
	score := int(math.Pow(2, float64(matches-1)))
	//fmt.Println(c.winning, c.roster, matches, score)
	return score
}

func (p *Pile) solve1() int {
	sum := 0
	for _, c := range p.cards {
		sum += c.score()
	}
	return sum
}

func (p *Pile) solve2() int {
	dupes := []*Card{}
	for _, c := range p.cards {
		dupes = append(dupes, c)
	}

	i := 0
	for {
		card := dupes[i]
		matches := card.matches()
		for j := 0; j < matches; j++ {
			dupes = append(dupes, p.cards[card.number+j])
		}

		i++
		if i == len(dupes) {
			break
		}
	}

	//fmt.Println(dupes)
	return len(dupes)
}

func (c *Card) String() string {
	return fmt.Sprintf("(#%d)", c.number)
}
