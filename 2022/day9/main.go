package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Left      string = "L"
	Right     string = "R"
	Up        string = "U"
	Down      string = "D"
	UpLeft    string = "UL"
	UpRight   string = "UR"
	DownLeft  string = "DL"
	DownRight string = "DR"
)

type Move struct {
	direction string
	distance  int
}

type Point struct {
	x int
	y int
}

type SnakePart struct {
	path []*Point
}

type Snake struct {
	head SnakePart
	tail SnakePart
}

func main() {
	inputFileName := "input.txt"
	data, _ := os.ReadFile(inputFileName)

	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]

	head := SnakePart{path: []*Point{}}
	tail := SnakePart{path: []*Point{}}
	snake := Snake{head: head, tail: tail}

	for _, line := range lines {
		splat := strings.Split(line, " ")
		distance, _ := strconv.Atoi(splat[1])
		move := Move{direction: splat[0], distance: distance}
		snake.move(move)
	}

	fmt.Println("answer1", snake.tail.countUnique()) //6209
}

func (sp *SnakePart) position() *Point {
	if len(sp.path) == 0 {
		sp.path = append(sp.path, &Point{x: 0, y: 0})
	}
	return sp.path[len(sp.path)-1]
}

func (s *Snake) move(m Move) {
	//fmt.Println(m)
	for i := 0; i < m.distance; i++ {
		s.head.step(m.direction)
		direction := s.tail.follow(s.head)
		s.tail.step(direction)
	}
}

func (sp *SnakePart) step(direction string) {
	current := copy(sp.position())
	current.step(direction)
	sp.path = append(sp.path, current)
}

func copy(p *Point) *Point {
	return &Point{x: p.x, y: p.y}
}

func (p *Point) step(direction string) {
	if direction == Up {
		p.y = p.y - 1
	}
	if direction == Down {
		p.y = p.y + 1
	}
	if direction == Left {
		p.x = p.x - 1
	}
	if direction == Right {
		p.x = p.x + 1
	}

	if direction == UpLeft {
		p.step(Left)
		p.step(Up)
	}
	if direction == UpRight {
		p.step(Right)
		p.step(Up)
	}
	if direction == DownLeft {
		p.step(Left)
		p.step(Down)
	}
	if direction == DownRight {
		p.step(Right)
		p.step(Down)
	}
}

func (p Point) isLeft(t Point) bool {
	return p.x < t.x
}
func (p Point) isRight(t Point) bool {
	return p.x > t.x
}
func (p Point) isUp(t Point) bool {
	return p.y < t.y
}
func (p Point) isDown(t Point) bool {
	return p.y > t.y
}

func (p Point) isTouching(t Point) bool {
	return p.x >= t.x-1 && p.x <= t.x+1 && p.y >= t.y-1 && p.y <= t.y+1
}

func (sp *SnakePart) follow(target SnakePart) string {
	cur := sp.position()
	pos := target.position()

	if pos.isTouching(*cur) {
		return "N"
	}

	if pos.x == cur.x {
		if pos.isUp(*cur) {
			return Up
		}
		if pos.isDown(*cur) {
			return Down
		}
	}
	if pos.y == cur.y {
		if pos.isLeft(*cur) {
			return Left
		}
		if pos.isRight(*cur) {
			return Right
		}
	}

	if pos.isUp(*cur) && pos.isLeft(*cur) {
		return UpLeft
	}
	if pos.isUp(*cur) && pos.isRight(*cur) {
		return UpRight
	}
	if pos.isDown(*cur) && pos.isRight(*cur) {
		return DownRight
	}
	if pos.isDown(*cur) && pos.isLeft(*cur) {
		return DownLeft
	}

	panic("?")
}

func (p *Point) uniqueString() string {
	return fmt.Sprint(p.x) + ":" + fmt.Sprint(p.y)
}

func (sp *SnakePart) countUnique() int {
	seen := make(map[string]bool)
	for _, p := range sp.path {
		seen[p.uniqueString()] = true
	}
	return len(seen)
}
