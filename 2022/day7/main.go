package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type file struct {
	isDir    bool
	size     int
	name     string
	parent   *file
	children []*file
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func createDir(name string) *file {
	file := file{name: name, size: -1, isDir: true, parent: nil, children: make([]*file, 0)}
	return &file
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)
	lines := strings.Split(string(data), "\n")

	root := *createDir("root")
	fmt.Println(root)
	current := &root

	for _, line := range lines {
		if strings.HasPrefix(line, "$ cd") {
			processCd(current, line)
		}
		if strings.HasPrefix(line, "dir") {
			processDir(current, line)
		}

		fmt.Println("TREE")
		printTree(current, 0)
		fmt.Println()
	}
}

func findChildren(current *file, name string) (*file, error) {
	for _, file := range current.children {
		if file.isDir && file.name == name {
			return file, nil
		}
	}
	return nil, errors.New("children not found")
}

func processCd(current *file, line string) {
	regex, _ := regexp.Compile(`\$ cd ([a-z]+|\.\.)`)
	if !regex.MatchString(line) {
		return
	}

	matches := (regex.FindStringSubmatch(line))
	name := matches[1]

	if name == ".." {
		*current = *current.parent
		return
	}

	child, err := findChildren(current, name)
	check(err)
	*current = *child
}

func processDir(current *file, line string) {
	regex, _ := regexp.Compile("dir ([a-z]+)")
	matches := (regex.FindStringSubmatch(line))
	name := matches[1]

	newDir := createDir(name)
	newDir.parent = current

	current.children = append(current.children, newDir)
}

func printTree(current *file, deepness int) {
	for i := 0; i < deepness; i++ {
		fmt.Print("-")
	}
	fmt.Print(current.name)
	if current.parent != nil {
		fmt.Print(" (", current.parent.name, ")")
	}
	fmt.Println()
	for _, child := range current.children {
		printTree(child, deepness+1)
	}
}
