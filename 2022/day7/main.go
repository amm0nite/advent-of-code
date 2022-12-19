package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type file struct {
	isDir    bool
	size     int
	name     string
	parent   *file
	children []*file
}

type explorer struct {
	current *file
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func createDir(name string) *file {
	file := file{name: name, size: 0, isDir: true, parent: nil, children: make([]*file, 0)}
	return &file
}

func main() {
	inputFileName := "input.txt"
	data, err := os.ReadFile(inputFileName)
	check(err)
	lines := strings.Split(string(data), "\n")

	root := createDir("root")
	exp := &explorer{current: root}

	filePattern, _ := regexp.Compile(`^[0-9]+\s`)

	for _, line := range lines {
		if strings.HasPrefix(line, "$ cd") {
			processCd(exp, line)
		}
		if strings.HasPrefix(line, "dir") {
			processDir(exp, line)
		}
		if filePattern.MatchString(line) {
			processFile(exp, line)
		}
	}

	fmt.Println("TREE")
	printTree(root, 0)
	fmt.Println()

	answer1 := sumWithLimit(root, 100000)
	fmt.Println("answer1", answer1)

	totalSpace := 70000000
	spaceNeeded := 30000000
	spaceUsed := root.size
	spaceToFree := (spaceNeeded + spaceUsed) - totalSpace

	answer2 := findClosest(root, spaceToFree, 999999999)
	fmt.Println("answer2", answer2)
}

func findChildren(current *file, name string) (*file, error) {
	for _, file := range current.children {
		if file.isDir && file.name == name {
			return file, nil
		}
	}
	return nil, errors.New("children not found")
}

func processCd(exp *explorer, line string) {
	regex, _ := regexp.Compile(`\$ cd ([a-z]+|\.\.)`)
	if !regex.MatchString(line) {
		return
	}

	matches := (regex.FindStringSubmatch(line))
	name := matches[1]

	if name == ".." {
		exp.current = exp.current.parent
		return
	}

	child, err := findChildren(exp.current, name)
	check(err)
	exp.current = child
}

func processDir(exp *explorer, line string) {
	regex, _ := regexp.Compile("dir ([a-z]+)")
	matches := (regex.FindStringSubmatch(line))
	name := matches[1]

	newDir := createDir(name)
	newDir.parent = exp.current

	exp.current.children = append(exp.current.children, newDir)
}

func processFile(exp *explorer, line string) {
	regex, _ := regexp.Compile(`^([0-9]+)\s(.+)`)
	matches := (regex.FindStringSubmatch(line))
	size, _ := strconv.Atoi(matches[1])
	upwardSizeUpdate(exp.current, size)
}

func upwardSizeUpdate(current *file, size int) {
	current.size += size
	if current.parent != nil {
		upwardSizeUpdate(current.parent, size)
	}
}

func sumWithLimit(current *file, limit int) int {
	sum := 0
	if current.parent != nil && current.isDir && current.size <= limit {
		sum += current.size
	}
	for _, file := range current.children {
		sum += sumWithLimit(file, limit)
	}
	return sum
}

func findClosest(current *file, target int, best int) int {
	distance := current.size - target
	if distance > 0 && distance < (best-target) {
		best = current.size
	}
	for _, file := range current.children {
		best = findClosest(file, target, best)
	}
	return best
}

func printTree(current *file, deepness int) {
	for i := 0; i < deepness; i++ {
		fmt.Print("-")
	}
	fmt.Println(current.name, "(", current.size, ")")
	for _, child := range current.children {
		printTree(child, deepness+1)
	}
}
