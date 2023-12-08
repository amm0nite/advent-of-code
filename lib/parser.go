package lib

import (
	"strconv"
	"strings"
)

type Cursor struct {
	sentence int
	token    int
}

type Parser struct {
	Symbols   []rune
	Sentences []Sentence
	cursor    Cursor
}

type Token struct {
	Line   int
	Col    int
	Buffer string
}

type Sentence struct {
	Tokens []Token
}

func (t *Token) IsInt() bool {
	_, err := strconv.Atoi(t.Buffer)
	return err == nil
}

func (t *Token) ToInt() int {
	value, err := strconv.Atoi(t.Buffer)
	if err != nil {
		return 0
	}
	return value
}

func (t *Token) ToRune() rune {
	return rune(t.Buffer[0])
}

func (t *Token) Len() int {
	return len(t.Buffer)
}

func (t *Token) isEmpty() bool {
	return t.Buffer == ""
}

func (t *Token) append(c rune) {
	t.Buffer += string(c)
}

func (p *Parser) Next(validator func(t *Token) bool) *Token {
	for p.cursor.sentence < len(p.Sentences) {
		i := p.cursor.sentence

		for p.cursor.token < len(p.Sentences[i].Tokens) {
			j := p.cursor.token
			p.cursor.token++

			curr := p.Sentences[i].Tokens[j]
			if validator(&curr) {
				return &curr
			}
		}

		p.cursor.token = 0
		p.cursor.sentence++
	}
	return nil
}

func (p *Parser) ResetCursor() {
	p.cursor.sentence = 0
	p.cursor.token = 0
}

func (p *Parser) isSymbol(c rune) bool {
	for _, s := range p.Symbols {
		if c == s {
			return true
		}
	}
	return false
}

func (p *Parser) parse(pos int, line string) *Sentence {
	tokens := []Token{}
	current := &Token{Line: pos, Col: 1}

	for index, c := range line {
		col := index + 1

		if p.isSymbol(c) {
			if !current.isEmpty() {
				tokens = append(tokens, *current)
			}
			symbol := &Token{Line: pos, Col: col}
			symbol.append(c)
			tokens = append(tokens, *symbol)
			current = &Token{Line: pos, Col: col + 1}

			continue
		}

		current.append(c)
	}

	if !current.isEmpty() {
		tokens = append(tokens, *current)
	}

	return &Sentence{Tokens: tokens}
}

func CreateParser(input string, symbols []rune) *Parser {
	parser := &Parser{Symbols: symbols}

	lines := strings.Split(input, "\n")
	for index, line := range lines {
		if line == "" {
			continue
		}

		sentence := parser.parse(index+1, line)
		parser.Sentences = append(parser.Sentences, *sentence)
	}

	return parser
}
