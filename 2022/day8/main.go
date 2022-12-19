package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type World struct {
	forest [][]int
	width  int
	height int
}

func main() {
	inputFileName := "test_input.txt"
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

	answer1 := 0
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if world.isVisible(i, j) {
				fmt.Print(world.forest[i][j])
				answer1++
			} else {
				fmt.Print("x")
			}
		}
		fmt.Println()
	}
	fmt.Println("answer1", answer1)
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

		if tree < highest {
			return false
		}
		if _x == x && _y == y {
			return tree > highest
		}
		highest = tree
	}
	return true
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
