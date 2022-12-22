package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type World struct {
	forest [][]int
	width  int
	height int
}

func main() {
	inputFileName := "input.txt"
	data, _ := os.ReadFile(inputFileName)

	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]
	//fmt.Println(lines)

	width := len(lines[0])
	height := len(lines)
	//fmt.Println(width, height)

	forest := make([][]int, width)
	for i := range forest {
		forest[i] = make([]int, height)
	}

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			forest[i][j], _ = strconv.Atoi(string(lines[i][j]))
		}
	}

	world := World{forest: forest, width: width, height: height}

	visible := color.New(color.FgWhite)
	notVisible := color.New(color.FgRed)

	answer1 := 0
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if world.isVisible(i, j) {
				visible.Print(fmt.Sprint(world.forest[i][j]))
				answer1++
			} else {
				notVisible.Print(fmt.Sprint(world.forest[i][j]))
			}
		}
		fmt.Println()
	}

	fmt.Println("answer1", answer1) //1684

	answer2 := 0
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			score := world.sceneScore(i, j)
			if score > answer2 {
				answer2 = score
			}
		}
	}

	fmt.Println("answer2", answer2)
}

func (w World) look(x int, y int, orientation bool, direction bool) bool {
	highest := -1
	limit := 0

	if orientation {
		limit = w.height
	} else {
		limit = w.width
	}

	for k := 0; k < limit; k++ {
		_x, _y := 0, 0
		if orientation && direction {
			_x = x
			_y = k
		}
		if orientation && !direction {
			_x = x
			_y = (w.height - 1) - k
		}
		if !orientation && direction {
			_x = k
			_y = y
		}
		if !orientation && !direction {
			_x = (w.width - 1) - k
			_y = y
		}

		tree := w.forest[_x][_y]
		//fmt.Println("looking", orientation, direction, "(", _x, _y, ")", tree, highest)

		if _x == x && _y == y {
			return tree > highest
		}
		if tree > highest {
			highest = tree
		}
	}
	return true
}

func (w World) sceneLook(x int, y int, orientation bool, direction bool) int {
	highest := w.forest[x][y]

	_x := x
	_y := y

	counter := -1
	start := true

	for _x >= 0 && _x < w.width && _y >= 0 && _y < w.height {
		tree := w.forest[_x][_y]
		counter++

		//fmt.Println("looking", orientation, direction, "(", _x, _y, ")", tree, highest)

		if !start && tree >= highest {
			break
		}

		if orientation && direction {
			_x++
		}
		if orientation && !direction {
			_x--
		}
		if !orientation && direction {
			_y++
		}
		if !orientation && !direction {
			_y--
		}

		start = false
	}

	return counter
}

func (w World) isVisible(x int, y int) bool {
	res := false

	res = w.look(x, y, true, true)
	if res {
		return true
	}

	res = w.look(x, y, true, false)
	if res {
		return true
	}

	res = w.look(x, y, false, true)
	if res {
		return true
	}

	res = w.look(x, y, false, false)
	return res
}

func (w World) sceneScore(x int, y int) int {
	score1 := w.sceneLook(x, y, true, true)
	score2 := w.sceneLook(x, y, true, false)
	score3 := w.sceneLook(x, y, false, true)
	score4 := w.sceneLook(x, y, false, false)
	//fmt.Println("score", score1, score2, score3, score4)
	return score1 * score2 * score3 * score4
}
